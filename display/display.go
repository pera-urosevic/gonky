package display

import (
	"os"

	"github.com/gdamore/tcell"
	"github.com/pera-urosevic/gonky/util"
)

var screen tcell.Screen
var width int
var height int

// Start //
func Start(w int, h int) {
	width = w
	h = height
	var e error
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	screen, e = tcell.NewScreen()
	util.ErrorLogExit(e)
	e = screen.Init()
	util.ErrorLogExit(e)
	screen.HideCursor()
	screen.EnableMouse()
	go run()
}

func run() {
	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC:
				Destroy()
				os.Exit(0)
			}
		case *tcell.EventMouse:
			if ev.Buttons() == tcell.Button2 {
				Destroy()
				os.Exit(0)
			}
		}
	}
}

// Sync //
func Sync() {
	screen.Sync()
}

// Destroy //
func Destroy() {
	screen.Fini()
}
