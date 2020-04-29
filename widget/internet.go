package widget

import (
	"math/rand"
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
	"github.com/sparrc/go-ping"
)

// Internet //
type Internet struct {
	Host    string
	Samples int
	Size    int
	Lag     float64
	Every   time.Duration
	State   *util.History
	Render  func(*Internet)
}

func (internet *Internet) id() string {
	if util.Debug {
		t := time.Now().Format("05")
		return t[0:1]
	}
	return time.Now().Format("Mon-15")
}

// Start //
func (internet Internet) Start() {
	internet.State = util.HistoryCreate(internet.Size, internet.id())
	go runner.Run(internet.Every, internet.tick)
}

func (internet *Internet) tick() {
	internet.sensor()
	internet.Render(internet)
}

func (internet *Internet) sensor() {
	if util.Debug {
		internet.State.Add(internet.id(), rand.Float64()*2)
		return
	}
	pinger, e := ping.NewPinger(internet.Host)
	if e != nil {
		internet.State.Add(internet.id(), 1000)
		util.ErrorLog(e)
		return
	}
	pinger.Count = internet.Samples
	pinger.Run()
	stats := pinger.Statistics()
	sensor := float64(stats.AvgRtt.Seconds())
	internet.State.Add(internet.id(), sensor)
}

// HistogramValues //
func (internet *Internet) HistogramValues() []int {
	var isLag = func(item float64) bool {
		return item > internet.Lag
	}

	histogram := []int{}
	for _, segment := range internet.State.Segments() {
		h := 0
		if segment.Contains(isLag) {
			h = 1
			if segment.Count(isLag) > segment.ItemsLength()/2 {
				h = 2
			}
		}
		histogram = append(histogram, h)
	}
	return histogram
}
