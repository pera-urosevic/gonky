package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/shirou/gopsutil/cpu"
)

// Processor //
type Processor struct {
	Every  time.Duration
	State  float64
	Render func(*Processor)
}

// Start //
func (processor Processor) Start() {
	go runner.Run(processor.Every, processor.tick)
}

func (processor *Processor) tick() {
	processor.sensor()
	processor.Render(processor)
}

func (processor *Processor) sensor() {
	sensor, e := cpu.Percent(0, false)
	if e != nil {
		return
	}
	processor.State = sensor[0]
}
