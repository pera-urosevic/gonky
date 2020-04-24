package render

import (
	"github.com/pera-urosevic/gonky/display"
)

// Bar //
func Bar(value float64, min float64, max float64, x int, y int, width int) {
	fraction := (value - min) / (max - min)
	fill := fraction * float64(width)
	bar := ""
	for x := 0; x < width; x++ {
		if float64(x) <= fill {
			bar += "▒"
		} else {
			bar += "░"
		}
	}
	display.Print(bar, x, y)
}
