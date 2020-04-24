package util

// SecondsToUnits //
func SecondsToUnits(seconds int) (int, int, int, int) {
	days := int(seconds / 60 / 60 / 24)
	seconds = seconds % (60 * 60 * 24)
	hours := int(seconds / 60 / 60)
	seconds = seconds % (60 * 60)
	minutes := int(seconds / 60)
	seconds = seconds % 60
	return days, hours, minutes, seconds
}
