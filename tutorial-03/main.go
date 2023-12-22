package main

import (
	"github.com/gdamore/tcell/v2"
	runewidth "github.com/mattn/go-runewidth"
)

func main() {
	// Open the screen
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()
	// Process events
	eventLoop(screen)
}

// Event loop
func eventLoop(screen tcell.Screen) {
	for {
		// Get the event
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// In case of a key event
			if ev.Key() == tcell.KeyEscape {
				// Exit if the ESC key is pressed
				return
			} else {
				otherKeyEvent(screen)
			}
		case *tcell.EventResize:
			// In case of a resize event
			screen.Sync()
		}
		// Update the screen
		screen.Show()
	}
}

var str []rune = []rune("こんにちは世界！")
var x, i = 0, 0

// Other event
func otherKeyEvent(screen tcell.Screen) {
	if i >= len(str) {
		x = 0
		i = 0
		screen.Clear()
	}
	screen.SetContent(x, 0, str[i], nil, tcell.StyleDefault)
	x += runewidth.RuneWidth(str[i])
	i++
}
