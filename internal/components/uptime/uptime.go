package uptime

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetUptime() (uint64, error) {
	file, err := os.Open("/proc/uptime")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return 0, scanner.Err()
	}

	fields := strings.Fields(scanner.Text())

	uptime, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	return uint64(uptime), nil
}
