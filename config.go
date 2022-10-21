package logwrapper

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var multiWriter zerolog.LevelWriter
var writers []io.Writer

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

func appendOutput(output io.Writer) {
	writers = append(writers, output)
	multiWriter = zerolog.MultiLevelWriter(writers...)
	logger = zerolog.New(multiWriter).With().Timestamp().Logger()

}

func setConsoleOutput(console bool) {
	if console {
		appendOutput(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

func setFileOutput(conf Config) error {
	if conf.File {
		file, err := os.OpenFile(fileName(conf.FileName, conf.FilePath), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}

		_, err = file.Stat()
		if err != nil {
			return err
		}
		appendOutput(file)
	}
	return nil
}

func fileName(filename string, filepath string) string {
	if len(filepath) > 0 {
		return fmt.Sprint(filepath, "/", filename)
	}
	return filename
}

func callerMarshalFunc(pc uintptr, file string, line int) string {
	dirs := strings.Split(file, "/")

	if len(dirs) > 4 {
		file = strings.Join(dirs[len(dirs)-4:], "/")
	}
	return file + ":" + strconv.Itoa(line)
}
