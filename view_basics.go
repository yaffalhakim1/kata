package main

import "github.com/egoist/mygo/ui"

// basicsView is the first thing a beginner sees: the handful of ideas that
// make Japanese make sense — the three scripts, that there is no word for
// "is", politeness, particles, word order. Tae Kim's guide teaches this
// way, by concept before vocabulary, and this page does the same.
func (a *app) basicsView(c *ui.Context) {
	t := c.Theme()
	ui.Scroll(c).Grow(1).Padding(pageMargin, 18).Gap(12).Children(func() {
		ui.Column(c).MaxWidth(contentMax).Gap(4).Children(func() {
			ui.Text(c, "Start here").FontSize(22).FontWeight(650)
			ui.Text(c, "Never studied Japanese? Read these six ideas first. "+
				"They are the ground everything else stands on.").
				FontSize(14).TextColor(t.TextMuted)
		})
		for i, concept := range basics() {
			a.conceptCard(c, i+1, concept)
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
