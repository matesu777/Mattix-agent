package collector

import (
	"github.com/matesu777/Mattix/internal/components"
	"github.com/matesu777/Mattix/internal/models"
)

type Collector struct {
	Metrics models.Metrics
}

func New() (*Collector, error) {
	network, err := components.NewNetwork()
	if err != nil {
		return nil, err
	}

	hostname, err := components.HostName()
	if err != nil {
		return nil, err
	}
	cpu := components.NewCpu()

	return &Collector{
		Metrics: models.Metrics{
			Hostname: hostname,
			CPU:      cpu,
			Memory:   &components.Memory{},
			Disk:     &components.Disk{},
			Network:  network,
		},
	}, nil
}

func (c *Collector) Update() error {

	if err := c.Metrics.CPU.Scan(); err != nil {
		return err
	}

	if err := c.Metrics.Memory.Scan(); err != nil {
		return err
	}

	if err := c.Metrics.Disk.Scan(); err != nil {
		return err
	}

	if err := c.Metrics.Network.Scan(); err != nil {
		return err
	}

	uptime, err := components.GetUptime()
	if err != nil {
		return err
	}

	c.Metrics.Uptime = uptime

	return nil
}
