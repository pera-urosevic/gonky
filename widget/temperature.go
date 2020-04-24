package widget

import (
	"strconv"
	"strings"
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
)

// Temperature //
type Temperature struct {
	Path   string
	Every  time.Duration
	State  float64
	Render func(*Temperature)
}

// Start //
func (temperature Temperature) Start() {
	go runner.Run(temperature.Every, temperature.tick)
}

func (temperature *Temperature) tick() {
	temperature.sensor()
	temperature.Render(temperature)
}

func (temperature *Temperature) sensor() {
	data, e := util.ReadFile(temperature.Path)
	if e != nil {
		util.ErrorLog(e)
		return
	}
	sensor, e := strconv.ParseFloat(strings.TrimSpace(data), 64)
	if e != nil {
		util.ErrorLog(e)
		return
	}
	sensor = sensor / 1000
	temperature.State = sensor
}
