package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Network struct {
	Name string `json:"name"`
	IPv4 string `json:"ipv4"`
	MAC  string `json:"mac"`

	RxBytes uint64 `json:"rx_bytes"`
	TxBytes uint64 `json:"tx_bytes"`

	RxSpeed uint64 `json:"rx_speed"`
	TxSpeed uint64 `json:"tx_speed"`

	prevRx uint64
	prevTx uint64
}

func NewNetwork() (Network, error) {
	iface, err := getDefaultInterface()
	if err != nil {
		return Network{}, err
	}

	network := Network{
		Name: iface.Name,
		MAC:  iface.HardwareAddr.String(),
	}

	addrs, err := iface.Addrs()
	if err == nil {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				if ip := ipNet.IP.To4(); ip != nil {
					network.IPv4 = ip.String()
					break
				}
			}
		}
	}

	return network, nil
}

func getDefaultInterface() (*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil || len(addrs) == 0 {
			continue
		}

		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				if ipNet.IP.To4() != nil {
					return &iface, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("nenhuma interface de rede ativa encontrada")
}

func (n *Network) Scan() error {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "Inter-") ||
			strings.HasPrefix(line, "face") {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		iface := strings.TrimSpace(parts[0])

		if iface != n.Name {
			continue
		}

		fields := strings.Fields(parts[1])

		if len(fields) < 16 {
			return fmt.Errorf("formato inválido em /proc/net/dev")
		}

		rx, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return err
		}

		tx, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return err
		}

		n.RxBytes = rx
		n.TxBytes = tx

		if n.prevRx != 0 {
			n.RxSpeed = rx - n.prevRx
		}

		if n.prevTx != 0 {
			n.TxSpeed = tx - n.prevTx
		}

		n.prevRx = rx
		n.prevTx = tx

		return nil
	}

	return scanner.Err()
}
