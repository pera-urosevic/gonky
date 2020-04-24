package widget

import (
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
	"github.com/shirou/gopsutil/host"
)

// Uptime //
type Uptime struct {
	Every time.Duration
	State struct {
		Days    int
		Hours   int
		Minutes int
		Seconds int
	}
	Render func(*Uptime)
}

// Start //
func (uptime Uptime) Start() {
	go runner.Run(uptime.Every, uptime.tick)
}

func (uptime *Uptime) tick() {
	uptime.sensor()
	uptime.Render(uptime)
}

func (uptime *Uptime) sensor() {
	sensor, e := host.Uptime()
	if e != nil {
		util.ErrorLog(e)
		return
	}
	uptime.State.Days, uptime.State.Hours, uptime.State.Minutes, uptime.State.Seconds = util.SecondsToUnits(int(sensor))
}
