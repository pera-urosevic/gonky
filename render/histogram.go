package render

// Histogram //
func Histogram(values []int, runes []rune, x int, y int) {
	histogram := []rune{}
	for i := 0; i < len(values); i++ {
		value := values[i]
		histogram = append(histogram, runes[value])
	}
	TextLeft(string(histogram), x, y)
}
