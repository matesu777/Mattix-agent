package components

import (
	"bufio"
	"github.com/matesu777/Mattix/internal/utils"
	"os"
	"runtime"
	"strings"
)

type CPU struct {
	Usage float64 `json:"usage"`
	Cores int     `json:"cores"`

	prevTotal uint64
	prevIdle  uint64
}

func NewCpu() *CPU {
	return &CPU{
		Cores: runtime.NumCPU(),
	}
}

func (c *CPU) Scan() error {
	data, err := os.Open("/proc/stat")
	if err != nil {
		return err
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	if !scanner.Scan() {
		return scanner.Err()
	}

	fields := strings.Fields(scanner.Text())

	user, err := utils.ConvertToUnit64(fields[1])
	if err != nil {
		return err
	}
	nice, err := utils.ConvertToUnit64(fields[2])
	if err != nil {
		return err
	}
	system, err := utils.ConvertToUnit64(fields[3])
	if err != nil {
		return err
	}
	idle, err := utils.ConvertToUnit64(fields[4])
	if err != nil {
		return err
	}
	iowait, err := utils.ConvertToUnit64(fields[5])
	if err != nil {
		return err
	}
	irq, err := utils.ConvertToUnit64(fields[6])
	if err != nil {
		return err
	}
	softirq, err := utils.ConvertToUnit64(fields[7])
	if err != nil {
		return err
	}
	steal, err := utils.ConvertToUnit64(fields[8])
	if err != nil {
		return err
	}

	idleTime := idle + iowait

	totalTime := user +
		nice +
		system +
		idle +
		iowait +
		irq +
		softirq +
		steal

	if c.prevTotal == 0 {
		c.prevTotal = totalTime
		c.prevIdle = idleTime
		return nil
	}

	totalDelta := totalTime - c.prevTotal
	idleDelta := idleTime - c.prevIdle

	if totalDelta > 0 {
		c.Usage = float64(totalDelta-idleDelta) / float64(totalDelta) * 100
	}

	c.prevTotal = totalTime
	c.prevIdle = idleTime

	return nil
}
