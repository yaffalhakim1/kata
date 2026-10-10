package main

import "github.com/egoist/mygo/ui"

// basicsView is the first thing a beginner sees: where to start, then the
// handful of ideas that make Japanese make sense — the three scripts, that
// there is no word for "is", politeness, particles, word order. Tae Kim's
// guide teaches this way, by concept before vocabulary, and this page does
// the same.
func (a *app) basicsView(c *ui.Context) {
	t := c.Theme()
	ui.Scroll(c).Grow(1).Padding(pageMargin, 18).Gap(12).Children(func() {
		ui.Column(c).MaxWidth(contentMax).Gap(4).Children(func() {
			ui.Text(c, "Start here").FontSize(22).FontWeight(650)
			ui.Text(c, "Never studied Japanese? Follow the four steps below, "+
				"in order. Then read the six ideas — they are the ground "+
				"everything else stands on.").
				FontSize(14).TextColor(t.TextMuted)
		})
		a.pathCard(c)
		for i, concept := range basics() {
			a.conceptCard(c, i+1, concept)
		}
	})
}

// pathStep is one stage of the learning path: where to go, and why.
type pathStep struct {
	N     int
	Title string
	Body  string
	Tab   int
	Cue   string // the button label, when the step has a page to go to
}

// learningPath is the order a beginner should move through the app. It is
// the same guide the README gives, kept here so it is in front of the
// learner rather than in a file they will not open.
func learningPath() []pathStep {
	return []pathStep{
		{1, "Read Basics", "The six ideas below, five minutes. Everything else assumes them.", TabBasics, ""},
		{2, "Learn the Kana", "Hiragana one row at a time, each with a picture. Read-only is enough.", TabKana, "Open Kana"},
		{3, "Study the Words", "Now the words make sense: meaning, reading, and how to use them.", TabStudy, "Open Words"},
		{4, "Quiz yourself", "Test what stuck. Every answer shows the word again.", TabQuiz, "Open Quiz"},
	}
}

// pathCard is the four-step guide, with a button on each step that jumps to
// the page it names, so a confused beginner is one click from the right
// place.
func (a *app) pathCard(c *ui.Context) {
	t := c.Theme()
	card(c).MaxWidth(contentMax).Children(func() {
		ui.Text(c, "Your path").FontSize(16).FontWeight(600)
		hint(c, "Already know hiragana? Skip to Words, or hit Quiz to see where you are.")
		for _, s := range learningPath() {
			ui.Row(c).AlignItems(ui.Center).Gap(12).Children(func() {
				ui.Textf(c, "%d", s.N).FontSize(13).FontWeight(700).
					TextColor(t.AccentText).Background(t.Accent).
					Radius(20).Padding(2, 8).Shrink(0)
				ui.Column(c).Grow(1).Shrink(1).MinWidth(0).Gap(2).Children(func() {
					ui.Text(c, s.Title).FontSize(14).FontWeight(600)
					ui.Text(c, s.Body).FontSize(12).TextColor(t.TextMuted)
				})
				if s.Cue != "" {
					ui.Button(c, s.Cue).Shrink(0).OnClick(func() { a.tab = s.Tab })
				}
			})
		}
	})
}

// conceptCard is one idea: a number, a title, the explanation, and an
// example with its reading.
func (a *app) conceptCard(c *ui.Context, n int, concept Concept) {
	t := c.Theme()
	card(c).MaxWidth(contentMax).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
			ui.Textf(c, "%d", n).FontSize(13).FontWeight(700).
				TextColor(t.AccentText).Background(t.Accent).
				Radius(20).Padding(2, 8).Shrink(0)
			ui.Text(c, concept.Title).FontSize(16).FontWeight(600)
		})
		ui.Text(c, concept.Body).FontSize(14).TextColor(t.Text)
		ui.Column(c).Gap(3).Padding(12, 14).Radius(10).
			Background(t.Background).Border(1, controlBorder(t.Dark)).Children(func() {
			ui.Text(c, concept.Example).FontSize(15)
			ui.Text(c, concept.Gloss).FontSize(12).TextColor(t.TextMuted)
		})
	})
}
