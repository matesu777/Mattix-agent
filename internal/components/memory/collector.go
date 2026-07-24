package memory

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func (m *Memory) Collector() error {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			total, err := strconv.ParseUint(fields[1], 10, 64)

			if err != nil {
				return err
			}
			m.Total = total * 1024
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			free, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return err
			}
			m.Free = free * 1024
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	m.Used = m.Total - m.Free

	return nil
}
