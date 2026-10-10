package main

import (
	"github.com/egoist/mygo/ui"
)

// studyView is the browse-and-learn page: a search, a topic filter, and the
// words themselves, each as a card that shows the English, the Indonesian
// and an example sentence.
func (a *app) studyView(c *ui.Context) {
	t := c.Theme()
	shown := a.studyCards()

	ui.Column(c).Grow(1).MinWidth(0).Children(func() {
		// The controls, on the window's margin. The search takes the room
		// the topic leaves, so neither clips at the smallest window.
		ui.Row(c).AlignItems(ui.Center).Gap(12).Padding(pageMargin, 14).Children(func() {
			ui.SearchField(c, &a.query).Label("Search words").Placeholder("Search a word…").
				Bind(&a.searchBox).Grow(1).Shrink(1).MinWidth(0)
			topics := append([]string{"All"}, topicNames(a.cards)...)
			ui.Select(c, &a.topic, topics).Label("Topic").Shrink(0)
		})
		ui.Divider(c).Background(t.Border)

		// The words, in a column that scrolls.
		if len(shown) == 0 {
			a.noWords(c)
			return
		}
		ui.Scroll(c).Grow(1).Padding(pageMargin, 16).Gap(10).Children(func() {
			for _, card := range shown {
				a.wordCard(c, card)
			}
		})
	})
}

// studyCards is the deck filtered by the topic and the search, in a stable
// order.
func (a *app) studyCards() []Card {
	cards := cardsIn(a.cards, a.topicOrEmpty())
	if q := a.query; q != "" {
		cards = filterCards(cards, q)
	}
	return cards
}

// topicOrEmpty turns the "All" choice into the empty filter cardsIn wants.
func (a *app) topicOrEmpty() string {
	if a.topic == "" || a.topic == "All" {
		return ""
	}
	return a.topic
}

// wordCard is one word: the Japanese, how it is read, the meaning, an
// example, and a note on how the word is actually used — so a beginner
// learns the concept, not just the gloss.
func (a *app) wordCard(c *ui.Context, w Card) {
	t := c.Theme()
	card(c).MaxWidth(contentMax).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
			ui.Text(c, w.Japanese).FontSize(20).FontWeight(600).Shrink(1)
			ui.Text(c, w.Romaji).FontSize(13).TextColor(t.TextMuted).Shrink(1)
			ui.Spacer(c)
			ui.Text(c, w.Indo).FontSize(16).TextColor(t.Accent).Shrink(1)
			ui.Badge(c, w.Topic)
		})
		ui.Text(c, w.Example).FontSize(13).TextColor(t.TextMuted)
		if w.Note != "" {
			ui.Row(c).AlignItems(ui.Start).Gap(6).Children(func() {
				ui.Icon(c, iconNote).Size(13, 13).TextColor(t.TextMuted).Shrink(0)
				ui.Text(c, w.Note).FontSize(12).TextColor(t.TextMuted)
			})
		}
	})
}

// noWords is the empty state: it names what was searched for and offers a
// way out, rather than only saying there is nothing.
func (a *app) noWords(c *ui.Context) {
	t := c.Theme()
	message, body := "No words here", "Nothing matches. Try another topic or search."
	if a.query != "" {
		message = "No words match “" + a.query + "”"
		body = "Clear the search to see every word."
	}
	ui.Column(c).Grow(1).Center().Gap(10).Padding(40).Children(func() {
		ui.Text(c, message).FontSize(15).FontWeight(600).TextColor(t.TextMuted)
		ui.Text(c, body).FontSize(13).TextColor(t.TextMuted).MaxWidth(420)
		if a.query != "" {
			ui.Button(c, "Clear search").OnClick(func() { a.query = "" })
		}
	})
}
