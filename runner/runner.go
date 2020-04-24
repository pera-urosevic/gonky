package runner

import "time"

// Run //
func Run(duration time.Duration, callback func()) {
	callback()
	RunAfter(duration, callback)
}

// RunAfter //
func RunAfter(duration time.Duration, callback func()) {
	t := time.NewTicker(duration)
	for {
		select {
		case <-t.C:
			callback()
		}
	}
}
