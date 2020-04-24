package render

import (
	"github.com/pera-urosevic/gonky/display"
)

// TextLeft //
func TextLeft(text string, x int, y int) {
	display.Print(text, x, y)
}

// TextRight //
func TextRight(text string, x int, y int) {
	display.PrintReverse(text, x, y)
}
