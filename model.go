package main

import "math/rand"

// qa is one question and its answer, ready to be asked. The quiz works on
// these, so the same engine drives a word quiz and a kana quiz.
type qa struct {
	Key    string // stable identity for the question
	Prompt string // what the learner reads
	Answer string // the right choice
	Reveal string // extra line shown after answering, to teach
}

// optionsPerQuestion is how many choices a question shows, when the pool
// allows it.
const optionsPerQuestion = 4

// quiz is one run through a set of questions. It is pure state: the view
// reads it, and Choose and Next move it forward. Nothing here touches the
// system, so a whole run can be driven from a test.
type quiz struct {
	items    []qa
	order    []int // indices into items, in the run's order
	idx      int
	score    int
	chosen   int // the option picked this question, -1 before an answer
	options  []int
	finished bool
}

// newQuiz builds a run over items, shuffled from seed so a run is
// repeatable when the seed is fixed. limit caps the number of questions
// (0 or more than the set means every question).
func newQuiz(items []qa, limit int, seed int64) *quiz {
	q := &quiz{items: items, chosen: -1}
	q.order = make([]int, len(items))
	for i := range q.order {
		q.order[i] = i
	}
	rng := rand.New(rand.NewSource(seed))
	rng.Shuffle(len(q.order), func(i, j int) { q.order[i], q.order[j] = q.order[j], q.order[i] })
	if limit > 0 && limit < len(q.order) {
		q.order = q.order[:limit]
	}
	q.prepare()
	return q
}

// prepare builds the choices for the question at idx: the right answer plus
// distinct wrong ones, in a stable order for this question.
func (q *quiz) prepare() {
	q.options = nil
	q.chosen = -1
	if q.finished || q.idx >= len(q.order) {
		return
	}
	correct := q.order[q.idx]
	q.options = append(q.options, correct)

	// Wrong answers are other items whose answer differs from the right
	// one, so two options never read the same.
	answer := q.items[correct].Answer
	wrong := make([]int, 0, len(q.items))
	for i := range q.items {
		if i != correct && q.items[i].Answer != answer {
			wrong = append(wrong, i)
		}
	}
	rng := rand.New(rand.NewSource(seedFor(q.idx)))
	rng.Shuffle(len(wrong), func(i, j int) { wrong[i], wrong[j] = wrong[j], wrong[i] })
	want := optionsPerQuestion - 1
	if want > len(wrong) {
		want = len(wrong)
	}
	q.options = append(q.options, wrong[:want]...)
	rng.Shuffle(len(q.options), func(i, j int) { q.options[i], q.options[j] = q.options[j], q.options[i] })
}

// seedFor mixes the question index so each question's options differ yet
// stay the same on every rebuild of the frame.
func seedFor(i int) int64 { return int64(i)*2654435761 + 12345 }

// Current is the question being asked now.
func (q *quiz) Current() qa { return q.items[q.order[q.idx]] }

// Prompt is what the learner reads.
func (q *quiz) Prompt() string { return q.Current().Prompt }

// Choices returns the option answers in the order shown.
func (q *quiz) Choices() []string {
	out := make([]string, len(q.options))
	for i, oi := range q.options {
		out[i] = q.items[oi].Answer
	}
	return out
}

// CorrectIndex is which choice is the right answer, -1 if none.
func (q *quiz) CorrectIndex() int {
	if len(q.options) == 0 {
		return -1
	}
	correct := q.order[q.idx]
	for i, oi := range q.options {
		if oi == correct {
			return i
		}
	}
	return -1
}

// Answered reports whether the current question has been answered.
func (q *quiz) Answered() bool { return q.chosen >= 0 }

// Chosen is the option picked, -1 before an answer.
func (q *quiz) Chosen() int { return q.chosen }

// Choose records the option picked, if the question is unanswered. It
// reports whether the choice was right.
func (q *quiz) Choose(i int) bool {
	if q.finished || q.Answered() || i < 0 || i >= len(q.options) {
		return false
	}
	q.chosen = i
	right := i == q.CorrectIndex()
	if right {
		q.score++
	}
	return right
}

// Next moves to the following question, or finishes the run. It is a no-op
// before the current question is answered, so the learner cannot skip.
func (q *quiz) Next() {
	if q.finished || !q.Answered() {
		return
	}
	q.idx++
	if q.idx >= len(q.order) {
		q.finished = true
		return
	}
	q.prepare()
}

// Finished reports whether the run is over.
func (q *quiz) Finished() bool { return q.finished }

// Total is the number of questions in the run.
func (q *quiz) Total() int { return len(q.order) }

// Number is the question being asked, 1-based.
func (q *quiz) Number() int { return q.idx + 1 }

// Score is how many were answered right so far.
func (q *quiz) Score() int { return q.score }

// Fraction is how far through the run the learner is, 0 to 1.
func (q *quiz) Fraction() float64 {
	if len(q.order) == 0 {
		return 1
	}
	if q.finished {
		return 1
	}
	return float64(q.idx) / float64(len(q.order))
}

// Grade is a one-word verdict on the finished run.
func (q *quiz) Grade() string {
	if q.Total() == 0 {
		return "Belum ada kartu"
	}
	ratio := float64(q.score) / float64(q.Total())
	switch {
	case ratio == 1:
		return "Sempurna!"
	case ratio >= 0.8:
		return "Bagus sekali!"
	case ratio >= 0.5:
		return "Lumayan, terus berlatih"
	default:
		return "Coba lagi, pelan-pelan"
	}
}

// Direction is which way a word quiz asks: the Japanese shown, the
// Indonesian to pick, or the other way around.
type Direction int

const (
	JapaneseToIndo Direction = iota
	IndoToJapanese
)

// Reading is the romaji of a card, shown under a Japanese prompt or answer
// so a beginner can read it.
func (d Direction) Reading(c Card) string { return c.Romaji }

// wordQuestions builds the questions for a word quiz in the chosen
// direction. The reveal line teaches: it shows both sides and the example.
func wordQuestions(cards []Card, dir Direction) []qa {
	out := make([]qa, 0, len(cards))
	for _, c := range cards {
		prompt, answer := c.Japanese, c.Indo
		if dir == IndoToJapanese {
			prompt, answer = c.Indo, c.Japanese
		}
		out = append(out, qa{
			Key:    c.Japanese + "|" + c.Indo,
			Prompt: prompt,
			Answer: answer,
			Reveal: c.Japanese + "（" + c.Romaji + "）— " + c.Indo,
		})
	}
	return out
}

// kanaQuestions builds the questions for a kana quiz: the character shown,
// the reading to pick.
func kanaQuestions(kana []Kana) []qa {
	out := make([]qa, 0, len(kana))
	for _, k := range kana {
		out = append(out, qa{
			Key:    k.Char,
			Prompt: k.Char,
			Answer: k.Romaji,
			Reveal: k.Char + " = " + k.Romaji + " · " + k.Mnemonic,
		})
	}
	return out
}
