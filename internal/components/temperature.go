package components

import (
	"fmt"
	"github.com/matesu777/Mattix/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

type Temperature struct {
	CpuTemp uint64 `json:"cputemp"`
}

func GetSensor() (string, error) {
	entries, err := os.ReadDir("/sys/class/hwmon")
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		dir := filepath.Join("/sys/class/hwmon", entry.Name())

		nameBytes, err := os.ReadFile(filepath.Join(dir, "name"))
		if err != nil {
			continue
		}

		name := strings.TrimSpace(string(nameBytes))

		if name == "coretemp" || name == "k10temp" {
			return dir, nil
		}
	}

	return "", fmt.Errorf("cpu sensor not found")
}

func (t *Temperature) GetTemperature() error {
	sensorDir, err := GetSensor()
	if err != nil {
		return err
	}

	for i := 1; ; i++ {
		labelPath := filepath.Join(sensorDir, fmt.Sprintf("temp%d_label", i))

		labelBytes, err := os.ReadFile(labelPath)
		if err != nil {
			break
		}

		if strings.TrimSpace(string(labelBytes)) != "Package id 0" {
			continue
		}

		inputPath := filepath.Join(sensorDir, fmt.Sprintf("temp%d_input", i))

		tempBytes, err := os.ReadFile(inputPath)
		if err != nil {
			return err
		}

		temp, err := utils.ConvertToUnit64(strings.TrimSpace(string(tempBytes)))
		if err != nil {
			return err
		}

		t.CpuTemp = temp
		return nil
	}

	return fmt.Errorf("package temperature not found")
}
