package main

import (
	"math/rand"
)

// Direction is which way a quiz asks: the Japanese word shown, the
// Indonesian one to pick, or the other way around.
type Direction int

const (
	JapaneseToIndo Direction = iota
	IndoToJapanese
)

// Prompt is the side the learner reads.
func (d Direction) Prompt(c Card) string {
	if d == IndoToJapanese {
		return c.Indo
	}
	return c.Japanese
}

// Answer is the side the learner must choose.
func (d Direction) Answer(c Card) string {
	if d == IndoToJapanese {
		return c.Japanese
	}
	return c.Indo
}

// Reading is the romaji of a card, shown under a Japanese prompt or answer
// so a beginner can read it.
func (d Direction) Reading(c Card) string { return c.Romaji }

// Options is how many choices a question shows, when the deck allows it.
const optionsPerQuestion = 4

// quiz is one run through a set of cards. It is pure state: the view reads
// it, and Choose and Next move it forward. Nothing here touches the system,
// so a whole run can be driven from a test.
type quiz struct {
	cards    []Card
	order    []int // indices into cards, in the run's order
	dir      Direction
	idx      int
	score    int
	chosen   int // the option picked this question, -1 before an answer
	options  []int
	finished bool
}

// newQuiz builds a run over cards, shuffled from seed so a run is
// repeatable when the seed is fixed. limit caps the number of questions
// (0 or more than the deck means every card).
func newQuiz(cards []Card, dir Direction, limit int, seed int64) *quiz {
	q := &quiz{cards: cards, dir: dir, chosen: -1}
	q.order = make([]int, len(cards))
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

// prepare builds the four choices for the question at idx: the right answer
// plus distinct wrong ones, in a stable order for this question.
func (q *quiz) prepare() {
	q.options = nil
	q.chosen = -1
	if q.finished || q.idx >= len(q.order) {
		return
	}
	correct := q.order[q.idx]
	q.options = append(q.options, correct)

	// Wrong answers are other cards whose answer differs from the right
	// one, so two options never read the same.
	answer := q.dir.Answer(q.cards[correct])
	wrong := make([]int, 0, len(q.cards))
	for i := range q.cards {
		if i != correct && q.dir.Answer(q.cards[i]) != answer {
			wrong = append(wrong, i)
		}
	}
	rng := rand.New(rand.NewSource(q.seedFor(q.idx)))
	rng.Shuffle(len(wrong), func(i, j int) { wrong[i], wrong[j] = wrong[j], wrong[i] })
	want := optionsPerQuestion - 1
	if want > len(wrong) {
		want = len(wrong)
	}
	q.options = append(q.options, wrong[:want]...)
	rng.Shuffle(len(q.options), func(i, j int) { q.options[i], q.options[j] = q.options[j], q.options[i] })
}

// seedFor mixes the run's index with the question, so each question's
// options differ yet stay the same on every rebuild of the frame.
func (q *quiz) seedFor(i int) int64 { return int64(i)*2654435761 + 12345 }

// Card is the question being asked now.
func (q *quiz) Card() Card { return q.cards[q.order[q.idx]] }

// Prompt is what the learner reads.
func (q *quiz) Prompt() string { return q.dir.Prompt(q.Card()) }

// Choices returns the options in the order shown.
func (q *quiz) Choices() []Card {
	out := make([]Card, len(q.options))
	for i, oi := range q.options {
		out[i] = q.cards[oi]
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
		return "No cards"
	}
	ratio := float64(q.score) / float64(q.Total())
	switch {
	case ratio == 1:
		return "Sempurna! (Perfect)"
	case ratio >= 0.8:
		return "Bagus! (Great)"
	case ratio >= 0.5:
		return "Lumayan (Not bad)"
	default:
		return "Coba lagi (Try again)"
	}
}
