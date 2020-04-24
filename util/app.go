package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

const logfile = "gonky.log"

// Debug //
var Debug = os.Getenv("DEBUG") == "true"

// ChooseDuration //
func ChooseDuration(condition bool, durationTrue time.Duration, durationFalse time.Duration) time.Duration {
	if condition {
		return durationTrue
	}
	return durationFalse
}

// Logger //
func Logger() {
	_ = os.Remove(logfile)
	f, e := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	log.SetOutput(f)
	Log("gonky starting...")
}

// Log //
func Log(message interface{}) {
	log.Println(string(fmt.Sprintln(message)))
}

// ErrorLog //
func ErrorLog(e error) {
	if e != nil {
		Log(e.Error())
	}
}

// ErrorLogExit //
func ErrorLogExit(e error) {
	if e != nil {
		ErrorLog(e)
		os.Exit(1)
	}
}
