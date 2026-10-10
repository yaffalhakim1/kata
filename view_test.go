package main

import (
	"image"
	"testing"

	"github.com/egoist/mygo/ui"
)

// newTestApp is an app showing the built-in deck, as a first run would.
func newTestApp() *app {
	return &app{cards: deck(), topic: "All", limitIdx: 1}
}

// TestStudyListsWords checks the Study page shows words, their reading and
// their meaning.
func TestStudyListsWords(t *testing.T) {
	a := newTestApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	for _, want := range []string{"こんにちは", "konnichiwa", "halo", "Kata"} {
		if !tt.HasText(want) {
			t.Errorf("no %q on screen; texts are %q", want, tt.Texts())
		}
	}
}

// TestSearchFilters checks typing in the search narrows the list, which is
// the point of the field.
func TestSearchFilters(t *testing.T) {
	a := newTestApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()

	tt.Click("Search words")
	tt.Type("mizu")
	if got := a.studyCards(); len(got) != 1 || got[0].Indo != "air" {
		t.Errorf("searching mizu gave %d cards, want just air", len(got))
	}
}

// TestTopicFilter checks the topic filter narrows the list.
func TestTopicFilter(t *testing.T) {
	a := newTestApp()
	a.topic = "Makanan"
	if got := a.studyCards(); len(got) != 10 {
		t.Errorf("the Makanan topic has %d cards, want 10", len(got))
	}
	for _, c := range a.studyCards() {
		if c.Topic != "Makanan" {
			t.Errorf("a card of another topic slipped in: %q", c.Topic)
		}
	}
}

// TestEmptySearchOffersAWayOut checks the empty state names the query and
// offers to clear it, rather than only saying there is nothing.
func TestEmptySearchOffersAWayOut(t *testing.T) {
	a := newTestApp()
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

// TestQuizFlow drives a whole run: start, answer every question, and land
// on the results screen with a full score.
func TestQuizFlow(t *testing.T) {
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
		q := a.q
		right := a.direction().Answer(q.Card())
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
	a := newTestApp()
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
