package logwrapper

import (
	"fmt"
	"io"

	"github.com/rs/zerolog"
)

// package wide logger variable, every log operation should use logger instance
var logger zerolog.Logger

// wrapper type
type logWrapper struct {
	prefix string
}

// Output get logger output
func (lw *logWrapper) Output() io.Writer {
	return multiWriter
}

// SetOutput sets logger output, appends existing writer
func (lw *logWrapper) SetOutput(w io.Writer) {
	appendOutput(w)
}

// Prefix get logger prefix
func (lw *logWrapper) Prefix() string {
	return lw.prefix
}

// SetPrefix set logger prefix, throws error not implemented
func (lw *logWrapper) SetPrefix(p string) {
	lw.prefix = p
	if e := logger.Error(); e.Enabled() {
		e.Caller().Msg("SetPrefix is not implemented")
	}
}

// SetHeader set logger header
func (lw *logWrapper) SetHeader(h string) {
	if e := logger.Error(); e.Enabled() {
		e.Msg("SetHeader is not implemented")
	}
}

// Print print log
func (lw *logWrapper) Print(i ...interface{}) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(fmt.Sprint(i...))
	}
}

// Printf print formated log
func (lw *logWrapper) Printf(format string, args ...interface{}) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(fmt.Sprintf(format, args...))
	}
}

// Debug log debug
func (lw *logWrapper) Debug(i ...interface{}) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(fmt.Sprint(i...))
	}
}

// Debugf log formated debug
func (lw *logWrapper) Debugf(format string, args ...interface{}) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(fmt.Sprintf(format, args...))
	}
}

// Info log info
func (lw *logWrapper) Info(i ...interface{}) {
	if e := logger.Info(); e.Enabled() {
		e.Msg(fmt.Sprint(i...))
	}
}

// Infof log foramted info
func (lw *logWrapper) Infof(format string, args ...interface{}) {
	if e := logger.Info(); e.Enabled() {
		e.Msg(fmt.Sprintf(format, args...))
	}
}

// Warn log warn
func (lw *logWrapper) Warn(i ...interface{}) {
	if e := logger.Warn(); e.Enabled() {
		e.Msg(fmt.Sprint(i...))
	}
}

// Warnf log formated warn
func (lw *logWrapper) Warnf(format string, args ...interface{}) {
	if e := logger.Warn(); e.Enabled() {
		e.Msg(fmt.Sprintf(format, args...))
	}
}

// Err log error if error is not nil
func (lw *logWrapper) Err(err error) {
	if e := logger.Error(); e.Enabled() && err != nil {
		e.Caller().Msg(err.Error())
	}
}

// Error log error
func (lw *logWrapper) Error(i ...interface{}) {
	if e := logger.Error(); e.Enabled() {
		e.Caller().Msg(fmt.Sprint(i...))
	}
}

// Errorf log error whith format
func (lw *logWrapper) Errorf(format string, args ...interface{}) {
	if e := logger.Error(); e.Enabled() {
		e.Caller().Msg(fmt.Sprintf(format, args...))
	}
}

// Fatal log fatal
func (lw *logWrapper) Fatal(i ...interface{}) {
	if e := logger.Fatal(); e.Enabled() {
		e.Caller().Msg(fmt.Sprint(i...))
	}
}

// Fatalf log fatal with format
func (lw *logWrapper) Fatalf(format string, args ...interface{}) {
	if e := logger.Fatal(); e.Enabled() {
		e.Caller().Msg(fmt.Sprintf(format, args...))
	}
}

// Panic log panic
func (lw *logWrapper) Panic(i ...interface{}) {
	if e := logger.Panic(); e.Enabled() {
		e.Caller().Msg(fmt.Sprint(i...))
	}
}

// Panicf log panic with format
func (lw *logWrapper) Panicf(format string, args ...interface{}) {
	if e := logger.Panic(); e.Enabled() {
		e.Caller().Msg(fmt.Sprintf(format, args...))
	}
}
