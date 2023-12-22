package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// 画面を開く
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	w, h := screen.Size()
	// イベントを処理する
	eventLoop(screen)
	screen.Fini()
	fmt.Println(w, h)
}

// イベントループ
func eventLoop(screen tcell.Screen) {
	for {
		// イベントを取得する
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			// キーイベントの場合
			if ev.Key() == tcell.KeyEscape {
				// ESCキーが押されたら終了する
				return
			} else {
				otherEvent(screen)
			}
		case *tcell.EventResize:
			// リサイズイベントの場合
			screen.Sync()
		}
		// 画面を更新する
		screen.Show()
	}
}

var str []rune = []rune("Hello World!")
var x = 0

// その他のイベント
func otherEvent(screen tcell.Screen) {
	if x >= len(str) {
		x = 0
		screen.Clear()
	}
	screen.SetContent(x, 0, str[x], nil, tcell.StyleDefault)
	x++
}
