package main

import "github.com/egoist/mygo/ui"

// kanaView teaches hiragana the way it is actually taught: one row at a
// time, the five vowel sounds first, each character with a picture to
// remember it by. A learner who cannot read kana cannot read anything else
// in the app, so this comes before the words.
func (a *app) kanaView(c *ui.Context) {
	t := c.Theme()
	ui.Scroll(c).Grow(1).Padding(pageMargin, 18).Gap(16).Children(func() {
		ui.Column(c).MaxWidth(contentMax).Gap(4).Children(func() {
			ui.Text(c, "Hiragana").FontSize(22).FontWeight(650)
			ui.Text(c, "The 46 characters that spell everything. Learn a row "+
				"at a time, left to right. The a-row sets the five vowel sounds "+
				"every other row is built from.").
				FontSize(14).TextColor(t.TextMuted)
		})
		for _, row := range kanaTable() {
			a.kanaRowCard(c, row)
		}
		ui.Text(c, "Once you can read these, try the Kana quiz.").
			FontSize(13).TextColor(t.TextMuted).MaxWidth(contentMax)
	})
}

// kanaRowCard is one column of the table: its name, and its characters.
func (a *app) kanaRowCard(c *ui.Context, row kanaRow) {
	t := c.Theme()
	card(c).MaxWidth(contentMax).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(8).Children(func() {
			ui.Text(c, row.Name).FontSize(13).FontWeight(700).
				TextColor(t.AccentText).Background(t.Accent).
				Radius(20).Padding(2, 8).Shrink(0)
			ui.Text(c, row.Name+"-row").FontSize(13).TextColor(t.TextMuted)
		})
		// The characters wrap, so a narrow window stacks them rather than
		// pushing the last one off the edge.
		ui.Row(c).Wrap().Gap(10).Children(func() {
			for _, k := range row.Kana {
				a.kanaTile(c, k)
			}
		})
	})
}

// kanaTile is one character: the kana, its reading, and the picture that
// makes it stick.
func (a *app) kanaTile(c *ui.Context, k Kana) {
	t := c.Theme()
	ui.Column(c).Width(120).Gap(4).Padding(10, 12).Radius(10).
		Background(t.Background).Border(1, controlBorder(t.Dark)).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(8).Children(func() {
			ui.Text(c, k.Char).FontSize(26)
			ui.Text(c, k.Romaji).FontSize(14).TextColor(t.Accent)
		})
		ui.Text(c, k.Mnemonic).FontSize(11).TextColor(t.TextMuted)
	})
}
