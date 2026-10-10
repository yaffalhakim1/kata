package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

func main() {
	s := loadSettings()
	a := &app{
		cards:      deck(),
		appearance: parseAppearance(s.Appearance),
		topic:      s.Topic,
		limitIdx:   limitIndex(s.Limit),
	}

	mygo.App.WhenReady(func() {
		win := mygo.NewWindow(mygo.WindowOptions{
			Title:     "Kata",
			Width:     820,
			Height:    660,
			MinWidth:  420,
			MinHeight: 420,
			// Opens where the user left it last time.
			StateKey: "main",
			// The window shows the interface MyGo draws, not a web page.
			Content: ui.View(a.view),
		})
		a.win = win
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
