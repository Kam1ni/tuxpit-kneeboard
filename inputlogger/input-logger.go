package inputlogger

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

type InputLoggerEventHandler func(deviceName string, button int)

type InputLogger struct {
	loggers []*logger
	handler InputLoggerEventHandler
}

func NewInputLogger(deviceNames []string, callback InputLoggerEventHandler) *InputLogger {
	result := InputLogger{
		handler: callback,
	}
	for _, name := range deviceNames {
		if name == "_KEYBOARD_" {
			logger := newKeyboardLogger(&result)
			result.loggers = append(result.loggers, logger)
			continue
		}

		logger := newLogger(&result, name)
		if logger != nil {
			result.loggers = append(result.loggers, logger)
		}
	}
	return &result
}

func NewAllInputsLogger(callback InputLoggerEventHandler) *InputLogger {
	result := InputLogger{
		handler: callback,
	}

	result.loggers = append(result.loggers, newKeyboardLogger(&result))

	devices, err := evdev.ListDevicePaths()
	if err != nil {
		panic("Failed to read /dev/input\n" + err.Error())
	}
	for _, item := range devices {
		isKeyboard, err := checkFileIsKeyboard(item.Path, true)
		if err != nil {
			continue
		}
		if isKeyboard {
			continue
		}
		logger := newLoggerFromEvdevDevice(&result, item)
		if logger == nil {
			continue
		}
		fmt.Println("Adding device", item.Name)
		result.loggers = append(result.loggers, logger)
	}

	return &result
}

func (i *InputLogger) Close() {
	for _, logger := range i.loggers {
		if logger.closed {
			continue
		}
		logger.close()
	}
}
