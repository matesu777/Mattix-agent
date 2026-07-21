package models

import (
	"time"

	"github.com/matesu777/Mattix/internal/components"
)

type Metrics struct {
	Hostname string    `json:"hostname"`
	Uptime   uint64    `json:"uptime"`
	UpdateAt time.Time `json:"updateAt"`

	CPU         components.CPU         `json:"cpu"`
	Memory      components.Memory      `json:"memory"`
	Disk        components.Disk        `json:"disk"`
	Network     components.Network     `json:"network"`
	Temperature components.Temperature `json:"temperature"`
}
