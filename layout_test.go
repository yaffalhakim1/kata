package main

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// The controls have to stay inside the window at the smallest size it
// allows. Before this was tested, the appearance button and the quiz panel
// were laid out past the right edge and were unreachable.
const (
	minWidth  = 420
	minHeight = 420
)

// TestHeaderStaysInsideAtMinimum checks the header controls do not clip at
// the narrowest window.
func TestHeaderStaysInsideAtMinimum(t *testing.T) {
	a := newTestApp()
	tt := ui.NewTester(a.view, minWidth, minHeight)
	tt.Frame()
	for _, label := range []string{"Words", "Kana", "Quiz", "Appearance: System"} {
		box, ok := tt.Find(label)
		if !ok {
			t.Fatalf("%q is missing at %dx%d", label, minWidth, minHeight)
		}
		if box.X+box.W > minWidth {
			t.Errorf("%q ends at %.0f, past the %d-wide window", label, box.X+box.W, minWidth)
		}
	}
}

// TestQuizPanelStaysInsideAtMinimum checks the quiz panel and its choices
// stay inside the narrowest window.
func TestQuizPanelStaysInsideAtMinimum(t *testing.T) {
	a := studyApp()
	a.tab = TabQuiz
	a.startQuiz()
	tt := ui.NewTester(a.view, minWidth, minHeight)
	tt.Frame()

	if _, ok := tt.Find("Question 1 of 10"); !ok {
		t.Fatal("the quiz question is missing at the minimum size")
	}
	for _, label := range a.q.Choices() {
		box, ok := tt.Find(label)
		if !ok {
			t.Fatalf("the choice %q is missing", label)
		}
		if box.X+box.W > minWidth {
			t.Errorf("the choice %q ends at %.0f, past the %d-wide window",
				label, box.X+box.W, minWidth)
		}
	}
}

// TestNoHorizontalOverflow checks the widest content, a long example
// sentence, does not push past the window at its smallest size.
func TestNoHorizontalOverflow(t *testing.T) {
	a := studyApp()
	a.query = "benkyou"
	tt := ui.NewTester(a.view, minWidth, minHeight)
	tt.Frame()

	box, ok := tt.Find("日本語を勉強します。")
	if !ok {
		t.Fatal("the example sentence is missing at the minimum size")
	}
	if box.X+box.W > minWidth {
		t.Errorf("the example ends at %.0f, past the %d-wide window", box.X+box.W, minWidth)
	}
}

// TestKanaTilesWrapAtMinimum checks the kana row wraps its tiles rather
// than pushing the last ones off the right edge at the narrowest window.
func TestKanaTilesWrapAtMinimum(t *testing.T) {
	a := newTestApp()
	a.tab = TabKana
	tt := ui.NewTester(a.view, minWidth, minHeight)
	tt.Frame()

	// The a-row has five tiles; if they did not wrap, the last (お) would
	// be past the edge.
	box, ok := tt.Find("お")
	if !ok {
		t.Fatal("the last kana of the a-row is missing at the minimum size")
	}
	if box.X+box.W > minWidth {
		t.Errorf("お ends at %.0f, past the %d-wide window: the row did not wrap",
			box.X+box.W, minWidth)
	}
}

// TestHeaderKeepsTheNameAtMinimum checks the wordmark is not crushed by the
// other header controls at the narrowest window.
func TestHeaderKeepsTheNameAtMinimum(t *testing.T) {
	a := newTestApp()
	tt := ui.NewTester(a.view, minWidth, minHeight)
	tt.Frame()
	if !tt.HasText("Kata") {
		t.Errorf("the wordmark is not shown at the minimum size; texts %q", tt.Texts())
	}
}

// TestPathButtonsStayInsideAtMinimum checks the guide's buttons, which sit
// at the end of a row, stay inside the narrowest window.
func TestPathButtonsStayInsideAtMinimum(t *testing.T) {
	a := newTestApp()
	a.tab = TabBasics
	tt := ui.NewTester(a.view, minWidth, 900)
	tt.Frame()
	for _, label := range []string{"Open Kana", "Open Words", "Open Quiz"} {
		box, ok := tt.Find(label)
		if !ok {
			t.Fatalf("%q is missing at %dx%d", label, minWidth, 900)
		}
		if box.X+box.W > minWidth {
			t.Errorf("%q ends at %.0f, past the %d-wide window", label, box.X+box.W, minWidth)
		}
	}
}
