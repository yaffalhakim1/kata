package main

import "testing"

// TestQuizScoring drives the quiz model without a window, so the scoring
// and the option building are checked on their own.
func TestQuizScoring(t *testing.T) {
	items := wordQuestions(deck(), JapaneseToIndo)
	q := newQuiz(items, 5, 1)
	if q.Total() != 5 {
		t.Fatalf("a 5-question run has %d questions", q.Total())
	}
	if len(q.Choices()) != optionsPerQuestion {
		t.Errorf("a question offers %d choices, want %d", len(q.Choices()), optionsPerQuestion)
	}

	for !q.Finished() {
		right := q.CorrectIndex()
		if right < 0 {
			t.Fatal("a question has no right answer")
		}
		if !q.Choose(right) {
			t.Error("choosing the right answer reported wrong")
		}
		q.Next()
	}
	if q.Score() != 5 {
		t.Errorf("five right answers scored %d", q.Score())
	}
	if q.Grade() != "Sempurna!" {
		t.Errorf("a perfect run graded %q", q.Grade())
	}
}

// TestQuizCannotAnswerTwice checks a second pick is ignored, so the score
// cannot be gamed.
func TestQuizCannotAnswerTwice(t *testing.T) {
	q := newQuiz(wordQuestions(deck(), JapaneseToIndo), 5, 2)
	right := q.CorrectIndex()
	q.Choose(right)
	if q.Choose((right + 1) % optionsPerQuestion) {
		t.Error("a second answer was accepted")
	}
	if q.Score() != 1 {
		t.Errorf("two answers scored %d, want 1", q.Score())
	}
}

// TestOptionsAreDistinct checks two choices never read the same, which
// would make a question ambiguous.
func TestOptionsAreDistinct(t *testing.T) {
	q := newQuiz(wordQuestions(deck(), JapaneseToIndo), 0, 3)
	for !q.Finished() {
		seen := map[string]bool{}
		for _, text := range q.Choices() {
			if seen[text] {
				t.Fatalf("the question %q repeats the choice %q", q.Prompt(), text)
			}
			seen[text] = true
		}
		q.Choose(q.CorrectIndex())
		q.Next()
	}
}

// TestWordQuestionsSwapSides checks the two directions ask opposite sides.
func TestWordQuestionsSwapSides(t *testing.T) {
	w := Card{Japanese: "水", Romaji: "mizu", Indo: "air"}
	jp := wordQuestions([]Card{w}, JapaneseToIndo)[0]
	if jp.Prompt != "水" || jp.Answer != "air" {
		t.Errorf("Japanese to Indonesian is the wrong way round: %+v", jp)
	}
	id := wordQuestions([]Card{w}, IndoToJapanese)[0]
	if id.Prompt != "air" || id.Answer != "水" {
		t.Errorf("Indonesian to Japanese is the wrong way round: %+v", id)
	}
}

// TestKanaQuestions checks a kana question asks the character and answers
// with its reading, and reveals the mnemonic.
func TestKanaQuestions(t *testing.T) {
	items := kanaQuestions(allKana())
	if len(items) != 46 {
		t.Errorf("the kana set has %d questions, want 46", len(items))
	}
	var a qa
	for _, it := range items {
		if it.Key == "あ" {
			a = it
		}
	}
	if a.Prompt != "あ" || a.Answer != "a" {
		t.Errorf("the あ question is %+v", a)
	}
	if a.Reveal == "" {
		t.Error("a kana question has no reveal line to teach with")
	}
}

// TestDeckIsComplete checks every card has the fields a beginner needs:
// a reading, a meaning, an example, and a usage note.
func TestDeckIsComplete(t *testing.T) {
	for _, c := range deck() {
		if c.Japanese == "" || c.Romaji == "" || c.Indo == "" {
			t.Errorf("card %+v is missing a core field", c)
		}
		if c.Example == "" {
			t.Errorf("%s has no example", c.Japanese)
		}
		if c.Note == "" {
			t.Errorf("%s has no usage note, which is the whole point", c.Japanese)
		}
	}
}
