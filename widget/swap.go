package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/shirou/gopsutil/mem"
)

// Swap //
type Swap struct {
	Every  time.Duration
	State  float64
	Render func(*Swap)
}

// Start //
func (w Swap) Start() {
	go runner.Run(w.Every, w.tick)
}

func (w *Swap) tick() {
	w.sensor()
	w.Render(w)
}

func (w *Swap) sensor() {
	sensor, e := mem.VirtualMemory()
	if e != nil {
		return
	}
	w.State = (1 - (float64(sensor.SwapFree) / float64(sensor.SwapTotal))) * 100.0
}
