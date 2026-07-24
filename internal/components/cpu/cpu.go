package cpu

import (
	"runtime"
	"strconv"
	"strings"
)

type Cpu struct {
	Usage     float64 `json:"usage"`
	Cores     []Core  `json:"cores"`
	prevTotal uint64
	prevIdle  uint64
}

type Core struct {
	ID        int     `json:"id"`
	Usage     float64 `json:"usage"`
	prevTotal uint64
	prevIdle  uint64
}

func NewCpu() Cpu {
	cpu := Cpu{}

	numCores := runtime.NumCPU()

	cpu.Cores = make([]Core, numCores)

	for i := range numCores {
		cpu.Cores[i] = Core{
			ID: i,
		}
	}

	return cpu
}

func (c *Cpu) Scan() error {
	stats, err := ParseCPUStat("/proc/stat")
	if err != nil {
		return err
	}

	for _, stat := range stats {
		if stat.Name == "cpu" {
			c.Usage = CalculateUsage(stat, &c.prevTotal, &c.prevIdle)
			continue
		}

		idStr := strings.TrimPrefix(stat.Name, "cpu")
		id, err := strconv.Atoi(idStr)

		if err != nil || id >= len(c.Cores) {
			continue
		}

		c.Cores[id].Usage = CalculateUsage(stat, &c.Cores[id].prevTotal, &c.Cores[id].prevIdle)
	}

	return nil
}

func CalculateUsage(stat CPUStat, prevTotal *uint64, prevIdle *uint64) float64 {
	total := stat.TotalTime()
	idle := stat.IdleTime()

	if *prevTotal == 0 {
		*prevTotal = total
		*prevIdle = idle
		return 0
	}

	totalDelta := total - *prevTotal
	idleDelta := idle - *prevIdle

	*prevTotal = total
	*prevIdle = idle

	if totalDelta == 0 {
		return 0
	}

	return float64(totalDelta-idleDelta) / float64(totalDelta) * 100 // Usage
}
