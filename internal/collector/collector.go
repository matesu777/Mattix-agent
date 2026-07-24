package collector

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/matesu777/Mattix/internal/components/cpu"
	"github.com/matesu777/Mattix/internal/components/disk"
	"github.com/matesu777/Mattix/internal/components/hostname"
	"github.com/matesu777/Mattix/internal/components/memory"
	"github.com/matesu777/Mattix/internal/components/network"
	"github.com/matesu777/Mattix/internal/components/temperature"
	"github.com/matesu777/Mattix/internal/components/uptime"
	"github.com/matesu777/Mattix/internal/models"
)

type Collector struct {
	mu sync.RWMutex

	Metrics models.Metrics
}

func NewCollector() (*Collector, error) {
	network, err := network.NewNetwork()
	if err != nil {
		return nil, fmt.Errorf("Fail to initialize network: %w\n", err)
	}
	hostname, err := hostname.Collector()
	if err != nil {
		return nil, fmt.Errorf("Fail to initialize hostname: %w\n", err)
	}

	return &Collector{
		Metrics: models.Metrics{
			Hostname:    hostname,
			CPU:         cpu.NewCpu(),
			Memory:      memory.Memory{},
			Disk:        disk.Disk{},
			Network:     network,
			Temperature: temperature.Temperature{},
		},
	}, nil
}

func (c *Collector) FastStart() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		if err := c.UpdateCpu(); err != nil {
			log.Println(err)
		}

		if err := c.UpdateNetwork(); err != nil {
			log.Println(err)
		}

		if err := c.UpdateUptime(); err != nil {
			log.Println(err)
		}

		c.UpdateAtUpdate()

		c.mu.Unlock()
	}
}

func (c *Collector) SlowStart() {

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	if err := c.UpdateDisk(); err != nil {
		log.Println(err)
	}
	if err := c.UpdateMemory(); err != nil {
		log.Println(err)
	}
	if err := c.UpdateTemperature(); err != nil {
		log.Println(err)
	}

	for range ticker.C {
		c.mu.Lock()

		if err := c.UpdateDisk(); err != nil {
			log.Println(err)
		}
		if err := c.UpdateMemory(); err != nil {
			log.Println(err)
		}
		if err := c.UpdateTemperature(); err != nil {
			log.Println(err)
		}

		c.mu.Unlock()
	}
}

func (c *Collector) UpdateCpu() error {
	if err := c.Metrics.CPU.Collect(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) UpdateMemory() error {
	if err := c.Metrics.Memory.Collector(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) UpdateNetwork() error {
	if err := c.Metrics.Network.Collector(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) UpdateDisk() error {
	if err := c.Metrics.Disk.Collect(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) UpdateAtUpdate() {
	c.Metrics.UpdateAt = time.Now()
}

func (c *Collector) UpdateUptime() error {
	uptime, err := uptime.Collector()
	if err != nil {
		return err
	}
	c.Metrics.Uptime = uptime
	return nil
}

func (c *Collector) UpdateTemperature() error {
	if err := c.Metrics.Temperature.Collector(); err != nil {
		return err
	}
	return nil
}

func (c *Collector) GetMetrics() models.Metrics {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Metrics
}
