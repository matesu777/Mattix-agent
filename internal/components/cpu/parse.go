package cpu

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseLine(fields []string) (CPUStat, error) {
	values := make([]uint64, 8)

	for i := 1; i <= 8; i++ {

		value, err := strconv.ParseUint(
			fields[i],
			10,
			64,
		)

		if err != nil {
			return CPUStat{}, err
		}

		values[i-1] = value
	}

	return CPUStat{
		Name:    fields[0],
		User:    values[0],
		Nice:    values[1],
		System:  values[2],
		Idle:    values[3],
		Iowait:  values[4],
		IRQ:     values[5],
		SoftIRQ: values[6],
		Steal:   values[7],
	}, nil
}

func ParseCPUStat(path string) ([]CPUStat, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var stats []CPUStat

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		fields := strings.Fields(scanner.Text())

		if len(fields) < 5 {
			continue
		}

		name := fields[0]

		if !strings.HasPrefix(name, "cpu") {
			break
		}

		stat, err := parseLine(fields)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	return stats, scanner.Err()
}
