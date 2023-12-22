package main

import (
	"github.com/gdamore/tcell/v2"
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

	str := []rune("Hello World!")
	for x := 0; x < len(str); x++ {
		screen.SetContent(x, 0, str[x], nil, tcell.StyleDefault)
	}
	screen.Show()

	for {
		// Wait for the next event
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// Exit the application
			if ev.Key() == tcell.KeyEscape {
				return
			}
		}
	}
}
