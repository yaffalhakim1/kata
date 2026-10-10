package main

import (
	"image"
	"testing"

	"github.com/egoist/mygo/ui"
)

// newTestApp is an app showing the built-in deck and kana, as a first run
// would.
func newTestApp() *app {
	return &app{cards: deck(), kana: allKana(), topic: "All", limitIdx: 1}
}

// studyApp is a new app on the Words page, which the word tests drive.
func studyApp() *app {
	a := newTestApp()
	a.tab = TabStudy
	return a
}

// TestBasicsTeachesFirst checks the first page a beginner sees explains the
// ideas, not just the words.
func TestBasicsTeachesFirst(t *testing.T) {
	a := newTestApp()
	a.tab = TabBasics
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	for _, want := range []string{"Start here", "Three scripts, one language", "Particles mark the job of a word"} {
		if !tt.HasText(want) {
			t.Errorf("the basics page is missing %q; texts are %q", want, tt.Texts())
		}
	}
}

// TestKanaTeachesRows checks the Kana page shows characters with their
// readings and mnemonics.
func TestKanaTeachesRows(t *testing.T) {
	a := newTestApp()
	a.tab = TabKana
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	for _, want := range []string{"Hiragana", "あ", "a", "か"} {
		if !tt.HasText(want) {
			t.Errorf("the kana page is missing %q", want)
		}
	}
}

// TestStudyListsWords checks the Study page shows words, their reading,
// meaning and usage note.
func TestStudyListsWords(t *testing.T) {
	a := studyApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	for _, want := range []string{"こんにちは", "konnichiwa", "halo", "Kata"} {
		if !tt.HasText(want) {
			t.Errorf("no %q on screen; texts are %q", want, tt.Texts())
		}
	}
}

// TestSearchFilters checks typing in the search narrows the list.
func TestSearchFilters(t *testing.T) {
	a := studyApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	tt.Click("Search words")
	tt.Type("mizu")
	if got := a.studyCards(); len(got) != 1 || got[0].Indo != "air" {
		t.Errorf("searching mizu gave %d cards, want just air", len(got))
	}
}

// TestEmptySearchOffersAWayOut checks the empty state names the query and
// offers to clear it.
func TestEmptySearchOffersAWayOut(t *testing.T) {
	a := studyApp()
	a.query = "zzzzz"
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	if !tt.HasText("No words match “zzzzz”") {
		t.Errorf("the empty state does not name the query; texts %q", tt.Texts())
	}
	if !tt.HasText("Clear search") {
		t.Error("the empty state offers no way out")
	}
}

// TestWordQuizFlow drives a whole word run: start, answer every question,
// and land on the results screen with a full score.
func TestWordQuizFlow(t *testing.T) {
	a := newTestApp()
	a.limitIdx = 0 // five questions
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	if err := tt.Click("Quiz"); err != nil {
		t.Fatalf("no Quiz page: %v", err)
	}
	tt.Frame()
	if err := tt.Click("Start quiz"); err != nil {
		t.Fatalf("no Start button: %v", err)
	}
	tt.Frame()
	if a.q == nil {
		t.Fatal("starting the quiz left no run")
	}

	for !a.q.Finished() {
		right := a.q.Choices()[a.q.CorrectIndex()]
		if err := tt.Click(right); err != nil {
			t.Fatalf("no option %q: %v", right, err)
		}
		tt.Frame()
		if !a.q.Answered() {
			t.Fatal("clicking an option did not answer the question")
		}
		label := "Next"
		if a.q.Number() == a.q.Total() {
			label = "See results"
		}
		if err := tt.Click(label); err != nil {
			t.Fatalf("no %q button: %v", label, err)
		}
		tt.Frame()
	}

	if a.q.Score() != a.q.Total() {
		t.Errorf("answering every question right scored %d of %d", a.q.Score(), a.q.Total())
	}
	if !tt.HasText("Results") {
		t.Errorf("the run did not end on the results screen; texts %q", tt.Texts())
	}
}

// TestKanaQuizFlow drives a kana run end to end.
func TestKanaQuizFlow(t *testing.T) {
	a := newTestApp()
	a.tab = TabQuiz
	a.sourceIdx = quizKana
	a.limitIdx = 0
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	if err := tt.Click("Start quiz"); err != nil {
		t.Fatalf("no Start button on the kana setup: %v", err)
	}
	tt.Frame()
	if a.q == nil || a.q.Total() != 5 {
		t.Fatalf("the kana run is not five questions: %+v", a.q)
	}
	right := a.q.Choices()[a.q.CorrectIndex()]
	if err := tt.Click(right); err != nil {
		t.Fatalf("no kana option %q: %v", right, err)
	}
	tt.Frame()
	if !a.q.Answered() {
		t.Fatal("clicking a kana option did not answer")
	}
}

// TestAppearanceMenuOffersEveryChoice checks the button opens a menu with
// all three appearances.
func TestAppearanceMenuOffersEveryChoice(t *testing.T) {
	a := newTestApp()
	a.appearance = AppearanceDark
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	if err := tt.Click("Appearance: Dark"); err != nil {
		t.Fatalf("the appearance button did not open a menu: %v", err)
	}
	items := tt.Menu()
	for _, want := range []string{"System", "Light", "Dark"} {
		if !contains(items, want) {
			t.Errorf("the menu does not offer %q; menu is %q", want, items)
		}
	}
}

// TestAppearanceChangesTheTheme checks the choice reaches the theme.
func TestAppearanceChangesTheTheme(t *testing.T) {
	for _, c := range []struct {
		appearance Appearance
		desktop    bool
		wantDark   bool
	}{
		{AppearanceSystem, false, false},
		{AppearanceSystem, true, true},
		{AppearanceLight, true, false},
		{AppearanceDark, false, true},
	} {
		a := newTestApp()
		a.appearance = c.appearance
		tt := ui.NewTester(a.view, 820, 660)
		tt.SetDark(c.desktop)
		tt.Frame()
		if got := tt.Image(); isDark(got) != c.wantDark {
			t.Errorf("appearance %v with desktop dark=%v rendered the wrong polarity",
				c.appearance, c.desktop)
		}
	}
}

// TestKeyboardReachesSearch checks F focuses the search field.
func TestKeyboardReachesSearch(t *testing.T) {
	a := studyApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()
	tt.Key(0, ui.KeyF)
	if !tt.Focused("Search words") {
		t.Error("F did not focus the search field")
	}
}

// contains reports whether a list holds a string.
func contains(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}

// isDark reports whether a frame is a dark one, by its corner pixel.
func isDark(img *image.RGBA) bool {
	r, g, b, _ := img.At(2, 2).RGBA()
	return int(r>>8)+int(g>>8)+int(b>>8) < 200
}
