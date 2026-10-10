package main

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"
)

// quizView is the test-yourself page. Before a run it offers the direction
// and how many questions; during a run it asks a word and four meanings;
// when the run ends it shows the score and what to do next.
func (a *app) quizView(c *ui.Context) {
	ui.Column(c).Grow(1).MinWidth(0).Center().Padding(pageMargin, 20).Children(func() {
		ui.Column(c).FillWidth().MaxWidth(contentMax).Gap(14).Children(func() {
			switch {
			case a.q == nil:
				a.quizSetup(c)
			case a.q.Finished():
				a.quizDone(c)
			default:
				a.quizQuestion(c)
			}
		})
	})
}

// quizSetup is the screen before a run: the two choices and a start button.
func (a *app) quizSetup(c *ui.Context) {
	t := c.Theme()
	card(c).Children(func() {
		ui.Text(c, "Test yourself").FontSize(17).FontWeight(600)
		hint(c, "Answer a few words, then see how you did.")

		sectionTitle(c, "Ask me")
		ui.Segmented(c, &a.dirIdx, "日本語 → Indonesia", "Indonesia → 日本語").Label("Direction")

		sectionTitle(c, "Questions")
		ui.Segmented(c, &a.limitIdx, limitLabels()...).Label("How many questions")

		ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
			ui.PrimaryButton(c, "Start quiz").OnClick(a.startQuiz)
			ui.Textf(c, "%d words in this set", len(a.studyCards())).
				FontSize(13).TextColor(t.TextMuted)
		})
	})
}

// quizQuestion is one question: what to read, and the four choices.
func (a *app) quizQuestion(c *ui.Context) {
	q := a.q
	t := c.Theme()
	card(c).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
			ui.Textf(c, "Question %d of %d", q.Number(), q.Total()).
				FontSize(12).TextColor(t.TextMuted).Shrink(0)
			ui.Spacer(c)
			ui.Textf(c, "Score %d", q.Score()).FontSize(12).TextColor(t.TextMuted).Shrink(0)
		})
		ui.Progress(c, q.Fraction())

		sectionTitle(c, a.promptLabel())
		ui.Text(c, q.Prompt()).FontSize(28).FontWeight(650)
		if a.direction() == JapaneseToIndo {
			ui.Text(c, q.Card().Romaji).FontSize(14).TextColor(t.TextMuted)
		}

		choices := q.Choices()
		for i, ch := range choices {
			a.optionButton(c, i, a.answerText(ch), q)
		}

		if q.Answered() {
			right := q.Chosen() == q.CorrectIndex()
			feedback(c, right, q.Card())
			a.nextButton(c, q)
		}
	})
}

// promptLabel names which language is being read, so the question is
// unambiguous.
func (a *app) promptLabel() string {
	if a.direction() == IndoToJapanese {
		return "Indonesian"
	}
	return "日本語"
}

// answerText is the side of a card the learner must choose.
func (a *app) answerText(w Card) string { return a.direction().Answer(w) }

// optionButton is one of the four choices. Before an answer it is a plain
// control; after, the right one is marked green and a wrong pick red, so
// the result is carried by an icon and a tint, not by colour alone.
func (a *app) optionButton(c *ui.Context, i int, text string, q *quiz) {
	t := c.Theme()
	b := ui.ButtonBase(c).Padding(12, 14).Radius(10).Grow(1).Focusable().
		Background(t.Background).Border(1, controlBorder(t.Dark)).
		OnClick(func() { q.Choose(i) })

	answered := q.Answered()
	right := i == q.CorrectIndex()
	picked := q.Chosen() == i

	switch {
	case answered && right:
		b.Background(t.Success.Alpha(0.14)).Border(1, t.Success)
	case answered && picked:
		b.Background(t.Danger.Alpha(0.14)).Border(1, t.Danger)
	case !answered && b.Hovered():
		b.Background(t.SurfaceHover)
	}
	b.Disabled(answered)

	b.Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
			ui.Text(c, fmt.Sprintf("%d", i+1)).FontSize(12).TextColor(t.TextMuted).Shrink(0)
			ui.Text(c, text).FontSize(15).Shrink(1)
			ui.Spacer(c)
			if answered && right {
				ui.Icon(c, iconCheck).Size(16, 16).TextColor(t.Success)
			} else if answered && picked {
				ui.Icon(c, iconCross).Size(16, 16).TextColor(t.Danger)
			}
		})
	})
}

// feedback says whether the answer was right, and shows the word with its
// example, so a wrong answer teaches rather than only scores.
func feedback(c *ui.Context, right bool, w Card) {
	t := c.Theme()
	colour, icon, word := t.Success, iconCheck, "Right"
	if !right {
		colour, icon, word = t.Danger, iconCross, "Not quite"
	}
	ui.Row(c).AlignItems(ui.Center).Gap(8).Children(func() {
		ui.Icon(c, icon).Size(15, 15).TextColor(colour)
		ui.Text(c, word).FontSize(14).FontWeight(600).TextColor(colour)
	})
	ui.Textf(c, "%s（%s）— %s", w.Japanese, w.Romaji, w.Indo).FontSize(14)
	ui.Text(c, w.Example).FontSize(13).TextColor(t.TextMuted)
}

// nextButton advances, or finishes, the run.
func (a *app) nextButton(c *ui.Context, q *quiz) {
	label := "Next"
	if q.Number() == q.Total() {
		label = "See results"
	}
	ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
		ui.PrimaryButton(c, label).OnClick(q.Next)
		ui.Text(c, "Press Enter").FontSize(12).TextColor(c.Theme().TextMuted)
	})
}

// quizDone is the end of a run: the score, a verdict, and what to do next.
func (a *app) quizDone(c *ui.Context) {
	q := a.q
	t := c.Theme()
	card(c).Children(func() {
		ui.Text(c, "Results").FontSize(12).TextColor(t.TextMuted)
		ui.Textf(c, "%d / %d", q.Score(), q.Total()).FontSize(32).FontWeight(700)
		ui.Text(c, q.Grade()).FontSize(15).TextColor(t.TextMuted)

		ui.Row(c).Gap(10).Children(func() {
			ui.PrimaryButton(c, "Try again").OnClick(func() { a.seed = time.Now().UnixNano(); a.startQuiz() })
			ui.Button(c, "New quiz").OnClick(func() { a.q = nil })
		})
		ui.Text(c, "Try again reshuffles the same set.").FontSize(12).TextColor(t.TextMuted)
	})
}

// startQuiz begins a run over the words the Study page currently shows, so
// the topic and search carry into the quiz.
func (a *app) startQuiz() {
	cards := a.studyCards()
	if len(cards) == 0 {
		return
	}
	if a.seed == 0 {
		a.seed = time.Now().UnixNano()
	}
	a.q = newQuiz(cards, a.direction(), a.limitCount(), a.seed)
}

// quizKeys wires the keyboard: the digits pick an answer, Enter starts the
// run or moves to the next question.
func (a *app) quizKeys(c *ui.Context) {
	if a.tab != TabQuiz {
		return
	}
	for i := 0; i < optionsPerQuestion; i++ {
		c.OnShortcut(0, ui.Key1+ui.Key(i), func() {
			if a.q != nil && !a.q.Finished() {
				a.q.Choose(i)
			}
		})
	}
	c.OnShortcut(0, ui.KeyEnter, func() {
		switch {
		case a.q == nil:
			a.startQuiz()
		case !a.q.Finished():
			a.q.Next()
		}
	})
}