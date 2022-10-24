package logwrapper

import (
	"errors"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func mock_config() Config {
	return Config{
		Level:           "trace",
		Console:         true,
		File:            true,
		FileName:        "tester.log",
		FilePath:        "",
		TimeFieldFormat: "2006-01-02T15:04:05.999999",
	}
}

func mock_timestampFunc() time.Time {
	return time.Date(2008, 1, 8, 17, 5, 05, 0, time.UTC)
}

func mock_callerMarshalFunc(pc uintptr, file string, line int) string {
	mockFile := "/Users/oto/Projects/microsena/logwrapper/loggerwrapper_test.go"
	return callerMarshalFunc(pc, mockFile, 73)
}

func mock_openFileFunc(name string, flag int, perm os.FileMode) (*os.File, error) {
	if name == "error" {
		return nil, errors.New("open file error")
	} else if name == "Stat error" {
		return nil, nil
	}

	return os.OpenFile(name, flag, perm)
}

func mock_loggerFatal() *zerolog.Event {
	return logger.WithLevel(zerolog.FatalLevel)
}

func mock_loggerPanic() *zerolog.Event {
	return logger.WithLevel(zerolog.PanicLevel)
}
