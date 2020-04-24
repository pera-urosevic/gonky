package widget

import (
	"os"
	"time"

	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
	"github.com/shirou/gopsutil/process"
)

// Top //
type Top struct {
	Every time.Duration
	State struct {
		CPU *Proc
		MEM *Proc
	}
	Render func(*Top)
}

// Proc //
type Proc struct {
	PID   int32
	Usage float64
	CMD   string
}

// Start //
func (w Top) Start() {
	go runner.Run(w.Every, w.tick)
}

func (w *Top) tick() {
	w.sensor()
	w.Render(w)
}

func (w *Top) sensor() {
	processes, e := process.Processes()
	if e != nil {
		util.ErrorLog(e)
		return
	}
	var cpu float64 = -1.0
	var mem float32 = -1.0
	pid := os.Getpid()
	for _, p := range processes {
		if int(p.Pid) == pid {
			continue
		}
		pcmd, e := p.Cmdline()
		if e != nil {
			continue
		}
		pcpu, e := p.CPUPercent()
		if e != nil {
			continue
		}
		if pcpu > cpu {
			cpu = pcpu
			w.State.CPU = &Proc{
				PID:   p.Pid,
				Usage: float64(pcpu),
				CMD:   pcmd,
			}
		}
		pmem, e := p.MemoryPercent()
		if e != nil {
			continue
		}
		if pmem > mem {
			mem = pmem
			w.State.MEM = &Proc{
				PID:   p.Pid,
				Usage: float64(pmem),
				CMD:   pcmd,
			}
		}
	}
}
