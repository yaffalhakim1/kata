package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// Kata teaches common Japanese and Indonesian words: a browsable list to
// study, and a quiz to test yourself. It follows the ports brand, so the
// two apps look like one family.

// app is the window's state. The view builds the interface from it on the
// main thread; nothing here runs in the background.
type app struct {
	win *mygo.Window

	cards []Card // the built-in deck, read once
	kana  []Kana // the built-in kana, read once

	query string // the study search
	topic string // "" or "All" means every topic
	tab   int    // TabBasics, TabStudy, TabKana or TabQuiz

	// searchBox is the search field, so the keyboard can reach it.
	searchBox ui.Handle

	// Quiz setup, kept between runs.
	sourceIdx int // quizSource: words or kana
	dirIdx    int // index into the two directions
	limitIdx  int // index into limitChoices
	seed      int64

	// q is the run in progress, or nil at the setup screen.
	q *quiz

	appearance Appearance
}

// The sources a quiz can draw from.
const (
	quizWords = iota
	quizKana
)

// The question counts a run offers; 0 means the whole filtered deck.
var limitChoices = []int{5, 10, 20, 0}

func limitLabels() []string { return []string{"5", "10", "20", "All"} }

// limitCount is how many questions the chosen limit asks.
func (a *app) limitCount() int {
	if a.limitIdx < 0 || a.limitIdx >= len(limitChoices) {
		return 10
	}
	return limitChoices[a.limitIdx]
}

// dir is the direction the setup chose.
func (a *app) direction() Direction {
	if a.dirIdx == 1 {
		return IndoToJapanese
	}
	return JapaneseToIndo
}

// limitIndex is the choice matching a stored count, for loading settings.
func limitIndex(count int) int {
	for i, n := range limitChoices {
		if n == count {
			return i
		}
	}
	return 1 // 10
}

// view builds the window from the app's state.
func (a *app) view(c *ui.Context) {
	a.applyTheme(c)
	t := c.Theme()

	a.quizKeys(c)
	a.shortcuts(c)

	ui.Column(c).Fill().Background(t.Background).Children(func() {
		a.header(c)
		ui.Divider(c).Background(t.Border)
		switch a.tab {
		case TabBasics:
			a.basicsView(c)
		case TabKana:
			a.kanaView(c)
		case TabQuiz:
			a.quizView(c)
		default:
			a.studyView(c)
		}
	})
}

// shortcuts wires the window's keys: F reaches the search field on the
// Study page.
func (a *app) shortcuts(c *ui.Context) {
	c.OnShortcut(0, ui.KeyF, func() { a.searchBox.Focus(c) })
}

// appearanceOf resolves the appearance to show: the user's choice, or what
// the desktop asks for when the choice is to follow it.
func (a *app) appearanceOf(c *ui.Context) bool {
	switch a.appearance {
	case AppearanceLight:
		return false
	case AppearanceDark:
		return true
	default:
		return c.Theme().Dark
	}
}

// applyTheme installs the theme for the current appearance, so a change
// takes effect on the next frame.
func (a *app) applyTheme(c *ui.Context) {
	c.SetTheme(theme(a.appearanceOf(c)))
}
