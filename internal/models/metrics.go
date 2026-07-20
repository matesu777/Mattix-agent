package models

import (
	"github.com/matesu777/Mattix/internal/components"
)

type Metrics struct {
	Hostname string `json:"hostname"`
	Uptime   uint64 `json:"uptime"`

	CPU     *components.CPU     `json:"cpu"`
	Memory  *components.Memory  `json:"memory"`
	Disk    *components.Disk    `json:"disk"`
	Network *components.Network `json:"network"`
}
