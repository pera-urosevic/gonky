package render

import (
	"fmt"
	"strconv"
	"strings"
)

// Paragraph //
func Paragraph(title string, x int, y int, width int, height int) {
	format := "%-" + strconv.Itoa(width) + "s"
	text := []string{}
	fields := strings.Fields(title)
	line := ""
	for _, field := range fields {
		if len(line)+len(field) < width {
			line += field + " "
		} else {
			text = append(text, fmt.Sprintf(format, line))
			line = field + " "
		}
	}
	text = append(text, fmt.Sprintf(format, line))
	if len(text) > height {
		text = text[:height]
	} else {
		for i := len(text); i < height; i++ {
			text = append(text, fmt.Sprintf(format, ""))
		}
	}

	for l, line := range text {
		TextLeft(line, x, y+l)
	}
}
