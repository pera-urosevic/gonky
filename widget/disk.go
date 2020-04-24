package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/shirou/gopsutil/disk"
)

// Disk //
type Disk struct {
	Path   string
	Every  time.Duration
	State  float64
	Render func(*Disk)
}

// Start //
func (w Disk) Start() {
	go runner.Run(w.Every, w.tick)
}

func (w *Disk) tick() {
	w.sensor()
	w.Render(w)
}

func (w *Disk) sensor() {
	sensor, e := disk.Usage(w.Path)
	if e != nil {
		return
	}
	w.State = sensor.UsedPercent
}
