package main

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestRender writes PNGs of the window, so a change to the interface can be
// looked at without launching the app.
func TestRender(t *testing.T) {
	a := newTestApp()
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()
	writePNG(t, "testdata/study.png", tt.Image())

	dark := newTestApp()
	ttd := ui.NewTester(dark.view, 820, 660)
	ttd.SetDark(true)
	ttd.Frame()
	writePNG(t, "testdata/study-dark.png", ttd.Image())
}

// TestRenderQuiz writes the quiz setup and a question in progress.
func TestRenderQuiz(t *testing.T) {
	a := newTestApp()
	a.tab = TabQuiz
	tt := ui.NewTester(a.view, 820, 660)
	tt.Frame()
	writePNG(t, "testdata/quiz-setup.png", tt.Image())

	a.startQuiz()
	tt.Frame()
	writePNG(t, "testdata/quiz-question.png", tt.Image())

	a.q.Choose(a.q.CorrectIndex())
	tt.Frame()
	writePNG(t, "testdata/quiz-answered.png", tt.Image())
}

func writePNG(t *testing.T, path string, img *image.RGBA) {
	t.Helper()
	if err := os.MkdirAll("testdata", 0o755); err != nil {
		t.Fatal(err)
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatal(err)
	}
}
