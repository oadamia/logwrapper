package mock

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	Level      string
	Console    bool
	FileName   string
	TimeFormat string
}

func (c Config) LogLevel() string {
	return c.Level
}

func (c Config) LogConsole() bool {
	return c.Console
}

func (c Config) LogFileName() string {
	return c.FileName
}

func (c Config) LogTimeFormat() string {
	return c.TimeFormat
}

func NewConfig() Config {
	return Config{
		Level:      "trace",
		Console:    true,
		FileName:   "tester.log",
		TimeFormat: "2006-01-02T15:04:05.999999",
	}
}

func TimestampFunc() time.Time {
	return time.Date(2008, 1, 8, 17, 5, 05, 0, time.UTC)
}

var RealCallerMarshalFunc func(uintptr, string, int) string

func CallerMarshalFunc(pc uintptr, file string, line int) string {

	mockFile := "/Users/oto/Projects/microsena/logwrapper/loggerwrapper_test.go"
	return RealCallerMarshalFunc(pc, mockFile, 73)

}

func OpenFileFunc(name string, flag int, perm os.FileMode) (*os.File, error) {
	if name == "error" {
		return nil, errors.New("open file error")
	} else if name == "Stat error" {
		return nil, nil
	}

	return os.OpenFile(name, flag, perm)
}
