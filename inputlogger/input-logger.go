package inputlogger

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

func (i *InputLogger) Close() {
	for _, logger := range i.loggers {
		logger.close()
	}
}
