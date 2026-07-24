package cpu

import (
	"runtime"
	"strconv"
	"strings"
)

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

func (c *Cpu) Collect() error {
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
