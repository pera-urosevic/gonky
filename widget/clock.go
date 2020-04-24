package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
)

// Clock //
type Clock struct {
	Every  time.Duration
	Render func(*Clock)
	State  time.Time
}

// Start //
func (clock Clock) Start() {
	go runner.Run(clock.Every, clock.tick)
}

func (clock *Clock) tick() {
	clock.sensor()
	clock.Render(clock)
}

func (clock *Clock) sensor() {
	clock.State = time.Now()
}
