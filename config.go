package logwrapper

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var multiWriter zerolog.LevelWriter
var writers []io.Writer
var openFileFunc = os.OpenFile

// Config wrapper configuration
type Config struct {
	Level           string
	File            bool
	FilePath        string
	FileName        string
	Console         bool
	TimeFieldFormat string
}

func configure(c Config) error {
	setLevel(c.Level)
	setTimeFieldFormat(c.TimeFieldFormat)
	setCallerMarshalFunction(callerMarshalFunc)
	setCallerSkipFrameCount(3)
	setTimestampFunc(utcTimeFunc)
	setConsoleOutput(c.Console)
	setTimestampFieldName("@timestamp")
	return setFileOutput(c)
}

func setLevel(str string) {
	level, _ := zerolog.ParseLevel(str)
	zerolog.SetGlobalLevel(level)
}

func setTimeFieldFormat(format string) {
	zerolog.TimeFieldFormat = format
}

func setTimestampFieldName(name string) {
	zerolog.TimestampFieldName = name
}

func setCallerSkipFrameCount(count int) {
	zerolog.CallerSkipFrameCount = count
}

func setCallerMarshalFunction(marshal func(pc uintptr, file string, line int) string) {
	zerolog.CallerMarshalFunc = marshal
}

func setTimestampFunc(timeFunc func() time.Time) {
	zerolog.TimestampFunc = timeFunc
}

func setOpenFileFunction(openFunc func(name string, flag int, perm os.FileMode) (*os.File, error)) {
	openFileFunc = openFunc
}

func setFileOutput(c Config) error {
	if c.File {
		file, err := openFile(c.FileName, c.FilePath)
		if err != nil {
			return err
		}
		appendOutput(file)
	}
	return nil
}

func setConsoleOutput(console bool) {
	if console {
		appendOutput(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

func appendOutput(output io.Writer) {
	writers = append(writers, output)
	multiWriter = zerolog.MultiLevelWriter(writers...)
	logger = zerolog.New(multiWriter).With().Timestamp().Logger()

}
