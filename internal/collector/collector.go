package collector

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/matesu777/Mattix/internal/components"
	"github.com/matesu777/Mattix/internal/models"
)

type Collector struct {
	mu sync.RWMutex

	Metrics models.Metrics
}

func New() (*Collector, error) {
	network, err := components.NewNetwork()
	if err != nil {
		return nil, fmt.Errorf("Fail to initialize network: %w\n", err)
	}
	hostname, err := components.HostName()
	if err != nil {
		return nil, fmt.Errorf("Fail to initialize hostname: %w\n", err)
	}

	return &Collector{
		Metrics: models.Metrics{
			Hostname:    hostname,
			CPU:         components.NewCpu(),
			Memory:      components.Memory{},
			Disk:        components.Disk{},
			Network:     network,
			Temperature: components.Temperature{},
		},
	}, nil
}

func (c *Collector) FastStart() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		if err := c.CpuUpdate(); err != nil {
			log.Println(err)
		}

		if err := c.NetworkUpdate(); err != nil {
			log.Println(err)
		}

		if err := c.UptimeUpdate(); err != nil {
			log.Println(err)
		}

		c.UpdateAtUpdate()

		c.mu.Unlock()
	}
}

func (c *Collector) SlowStart() {

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	if err := c.DiskUpdate(); err != nil {
		log.Println(err)
	}
	if err := c.MemoryUpdate(); err != nil {
		log.Println(err)
	}
	if err := c.TemperatureUpdate(); err != nil {
		log.Println(err)
	}

	for range ticker.C {
		c.mu.Lock()

		if err := c.DiskUpdate(); err != nil {
			log.Println(err)
		}
		if err := c.MemoryUpdate(); err != nil {
			log.Println(err)
		}

		if err := c.TemperatureUpdate(); err != nil {
			log.Println(err)
		}

		c.mu.Unlock()
	}
}

func (c *Collector) CpuUpdate() error {
	if err := c.Metrics.CPU.Scan(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) MemoryUpdate() error {
	if err := c.Metrics.Memory.Scan(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) NetworkUpdate() error {
	if err := c.Metrics.Network.Scan(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) DiskUpdate() error {
	if err := c.Metrics.Disk.Scan(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) UpdateAtUpdate() {
	c.Metrics.UpdateAt = time.Now()
}

func (c *Collector) UptimeUpdate() error {
	uptime, err := components.GetUptime()
	if err != nil {
		return err
	}
	c.Metrics.Uptime = uptime
	return nil
}

func (c *Collector) TemperatureUpdate() error {
	if err := c.Metrics.Temperature.GetTemperature(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) GetMetrics() models.Metrics {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Metrics
}
