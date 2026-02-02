package inputlogger

import (
	"fmt"
	"path"
	"strings"

	"github.com/MarinX/keylogger"
	"github.com/holoplot/go-evdev"
)

type logger struct {
	inputLogger     *InputLogger
	device          *evdev.InputDevice
	keyboardDevices []*keylogger.KeyLogger
	closed          bool
	name            string
}

func (l *logger) close() {
	l.closed = true
	if l.device != nil {
		l.device.Close()
	}
	for _, kb := range l.keyboardDevices {
		kb.Close()
	}
}

func (l *logger) run() {
	defer l.close()

	for true {
		if l.closed {
			return
		}
		event, err := l.device.ReadOne()
		if err != nil {
			panic("Failed to read event from " + l.name + "\n" + err.Error())
		}
		if event.Type != evdev.EV_KEY {
			continue
		}
		if event.Value != 1 {
			continue
		}
		l.inputLogger.handler(l.name, int(event.Code))
	}
}

func (l *logger) runKeyboard(keyboard *keylogger.KeyLogger, pth string) {
	defer l.close()

	for true {
		if l.closed {
			return
		}
		events := keyboard.Read()
		for event := range events {
			switch event.Type {
			case keylogger.EvKey:
				if event.KeyRelease() {
					l.inputLogger.handler("_KEYBOARD_", int(event.Code))
				}
			}
		}

	}
}

func newKeyboardLogger(inputLogger *InputLogger) *logger {
	// TODO: Ensure user is part of the input group
	fmt.Println("Creating new keyboard logger")
	result := logger{name: "_KEYBOARD_", inputLogger: inputLogger}
	allKeyboards, err := findKeyboards()
	if err != nil {
		panic("Failed to find keyboards\n" + err.Error())
	}
	for _, keyboard := range allKeyboards {
		kb, err := keylogger.New(path.Join("/dev/input", keyboard))
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "permission denied") {
				fmt.Printf("Can't open device %s\n%s\nSkipping...\n", keyboard, err.Error())
				continue
			} else {
				panic("Failed to open " + keyboard + "\n" + err.Error())
			}
		}

		result.keyboardDevices = append(result.keyboardDevices, kb)
		go result.runKeyboard(kb, keyboard)
	}
	return &result
}

func newLogger(inputLogger *InputLogger, deviceName string) *logger {
	devices, err := evdev.ListDevicePaths()
	if err != nil {
		panic("Failed to read /dev/input\n" + err.Error())
	}
	for _, item := range devices {
		if item.Name != deviceName {
			continue
		}

		logger := logger{
			name:        item.Name,
			inputLogger: inputLogger,
		}
		logger.device, err = evdev.Open(item.Path)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "permission denied") {
				fmt.Printf("Can't open device %s\n%s\nSkipping...\n", item.Name, err.Error())
				continue
			} else {
				panic("Failed to open " + item.Name + "\n" + err.Error())
			}
		}

		go logger.run()
		return &logger
	}
	return nil
}
