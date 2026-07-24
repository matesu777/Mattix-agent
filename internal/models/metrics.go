package models

import (
	"time"

	"github.com/matesu777/Mattix/internal/components/cpu"
	"github.com/matesu777/Mattix/internal/components/disk"
	"github.com/matesu777/Mattix/internal/components/memory"
	"github.com/matesu777/Mattix/internal/components/network"
	"github.com/matesu777/Mattix/internal/components/temperature"
)

type Metrics struct {
	Hostname string    `json:"hostname"`
	Uptime   uint64    `json:"uptime"`
	UpdateAt time.Time `json:"updateAt"`

	CPU         cpu.Cpu                 `json:"cpu"`
	Memory      memory.Memory           `json:"memory"`
	Disk        disk.Disk               `json:"disk"`
	Network     network.Network         `json:"network"`
	Temperature temperature.Temperature `json:"temperature"`
}
