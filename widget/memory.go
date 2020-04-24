package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/shirou/gopsutil/mem"
)

// Memory //
type Memory struct {
	Every  time.Duration
	State  float64
	Render func(*Memory)
}

// Start //
func (memory Memory) Start() {
	go runner.Run(memory.Every, memory.tick)
}

func (memory *Memory) tick() {
	memory.sensor()
	memory.Render(memory)
}

func (memory *Memory) sensor() {
	sensor, e := mem.VirtualMemory()
	if e != nil {
		return
	}
	memory.State = sensor.UsedPercent
}
