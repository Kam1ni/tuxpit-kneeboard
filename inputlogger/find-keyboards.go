package inputlogger

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

var validInputDeviceName = regexp.MustCompile(`^/dev/input/event\d+$`)

func findKeyboards() ([]string, error) {
	files, err := os.ReadDir("/dev/input")
	if err != nil {
		return nil, fmt.Errorf("Failed to read /dev/input\n%s", err.Error())
	}

	devices := []string{}
	for _, file := range files {
		fullPath := path.Join("/dev/input", file.Name())
		if !validInputDeviceName.MatchString(fullPath) {
			continue
		}
		cmd := exec.Command("udevadm", "info", "-q", "property", "-n", fullPath)
		buffer := bytes.NewBuffer(nil)
		cmd.Stdout = buffer
		err := cmd.Run()
		if err != nil {
			return nil, fmt.Errorf("Failed to run udevadm for device %s\n%s", fullPath, err.Error())
		}

		result := buffer.String()
		if strings.Contains(result, "\nID_INPUT_KEYBOARD=1\n") {
			devices = append(devices, file.Name())
		}
	}
	return devices, nil
}
