package display

import "github.com/gdamore/tcell"

// Print //
func Print(text string, x int, y int) {
	runes := []rune(text)
	for i, r := range runes {
		screen.SetContent(x+i, y, r, nil, tcell.StyleDefault)
	}
}

// PrintReverse //
func PrintReverse(text string, x int, y int) {
	runes := []rune(text)
	l := len(runes)
	for i := range runes {
		screen.SetContent(x-i, y, runes[l-i-1], nil, tcell.StyleDefault)
	}
}
