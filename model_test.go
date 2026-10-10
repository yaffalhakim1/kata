package main

import "testing"

// TestQuizScoring drives the quiz model without a window, so the scoring
// and the option building are checked on their own.
func TestQuizScoring(t *testing.T) {
	q := newQuiz(deck(), JapaneseToIndo, 5, 1)
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
	if q.Grade() != "Sempurna! (Perfect)" {
		t.Errorf("a perfect run graded %q", q.Grade())
	}
}

// TestQuizCannotAnswerTwice checks a second pick is ignored, so the score
// cannot be gamed.
func TestQuizCannotAnswerTwice(t *testing.T) {
	q := newQuiz(deck(), JapaneseToIndo, 5, 2)
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
	q := newQuiz(deck(), JapaneseToIndo, 0, 3)
	for !q.Finished() {
		seen := map[string]bool{}
		for _, c := range q.Choices() {
			text := q.dir.Answer(c)
			if seen[text] {
				t.Fatalf("the question %q repeats the choice %q", q.Prompt(), text)
			}
			seen[text] = true
		}
		q.Choose(q.CorrectIndex())
		q.Next()
	}
}

// TestDirectionSwapsSides checks the two directions ask opposite sides.
func TestDirectionSwapsSides(t *testing.T) {
	w := Card{Japanese: "水", Romaji: "mizu", Indo: "air"}
	if JapaneseToIndo.Prompt(w) != "水" || JapaneseToIndo.Answer(w) != "air" {
		t.Error("Japanese to Indonesian has the sides the wrong way")
	}
	if IndoToJapanese.Prompt(w) != "air" || IndoToJapanese.Answer(w) != "水" {
		t.Error("Indonesian to Japanese has the sides the wrong way")
	}
}
