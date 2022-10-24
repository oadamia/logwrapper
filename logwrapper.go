package logwrapper

import (
	"fmt"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

// package wide logger variable, every log operation should use logger instance
var logger zerolog.Logger

// fatal and panic functions, for testing purposes, variable dependency injection
var loggerFatal = logger.Fatal
var loggerPanic = logger.Panic

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

// Level get logger level
func (lw *logWrapper) Level() log.Lvl {
	return levelZeroToEcho(zerolog.GlobalLevel())
}

// SetLevel set logger level
func (lw *logWrapper) SetLevel(l log.Lvl) {
	zerolog.SetGlobalLevel(levelEchoToZero(l))
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

// Printj print json log
func (lw *logWrapper) Printj(j log.JSON) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(jsonMarshal(j))
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

// Debugj log json debug
func (lw *logWrapper) Debugj(j log.JSON) {
	if e := logger.Debug(); e.Enabled() {
		e.Msg(jsonMarshal(j))
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

// Infoj log json info
func (lw *logWrapper) Infoj(j log.JSON) {
	if e := logger.Info(); e.Enabled() {
		e.Msg(jsonMarshal(j))
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

// Warnj log json warn
func (lw *logWrapper) Warnj(j log.JSON) {
	if e := logger.Warn(); e.Enabled() {
		e.Msg(jsonMarshal(j))
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

// Errorj log error in json
func (lw *logWrapper) Errorj(j log.JSON) {
	if e := logger.Error(); e.Enabled() {
		e.Caller().Msg(jsonMarshal(j))
	}
}

// Fatal log fatal
func (lw *logWrapper) Fatal(i ...interface{}) {
	if e := loggerFatal(); e.Enabled() {
		e.Caller().Msg(fmt.Sprint(i...))
	}
}

// Fatalf log fatal with format
func (lw *logWrapper) Fatalf(format string, args ...interface{}) {
	if e := loggerFatal(); e.Enabled() {
		e.Caller().Msg(fmt.Sprintf(format, args...))
	}
}

// Fatalj log fatal in json
func (lw *logWrapper) Fatalj(j log.JSON) {
	if e := loggerFatal(); e.Enabled() {
		e.Caller().Msg(jsonMarshal(j))
	}
}

// Panic log panic
func (lw *logWrapper) Panic(i ...interface{}) {
	if e := loggerPanic(); e.Enabled() {
		e.Caller().Msg(fmt.Sprint(i...))
	}
}

// Panicf log panic with format
func (lw *logWrapper) Panicf(format string, args ...interface{}) {
	if e := loggerPanic(); e.Enabled() {
		e.Caller().Msg(fmt.Sprintf(format, args...))
	}
}

// Panicj log panic in json
func (lw *logWrapper) Panicj(j log.JSON) {
	if e := loggerPanic(); e.Enabled() {
		e.Caller().Msg(jsonMarshal(j))
	}
}
