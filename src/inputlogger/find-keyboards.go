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
		isKeyboard, err := checkFileIsKeyboard(fullPath, false)
		if err != nil {
			return nil, err
		}
		if isKeyboard {
			fmt.Println("Adding keyboard", file.Name())
			devices = append(devices, file.Name())
		}
	}
	return devices, nil
}

func checkFileIsKeyboard(fullPath string, includeMice bool) (bool, error) {
	if !validInputDeviceName.MatchString(fullPath) {
		return false, nil
	}
	cmd := exec.Command("udevadm", "info", "-q", "property", "-n", fullPath)
	buffer := bytes.NewBuffer(nil)
	cmd.Stdout = buffer
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("Failed to run udevadm for device %s\n%s", fullPath, err.Error())
	}

	result := buffer.String()
	isKeyboard := strings.Contains(result, "\nID_INPUT_KEYBOARD=1\n")
	if !isKeyboard && includeMice {
		isKeyboard = strings.Contains(result, "\nID_INPUT_MOUSE=1\n")
	}
	return isKeyboard, nil
}
