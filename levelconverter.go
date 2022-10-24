package logwrapper

import (
	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

func levelZeroToEcho(l zerolog.Level) log.Lvl {
	switch l {
	case zerolog.DebugLevel:
		return log.DEBUG
	case zerolog.WarnLevel:
		return log.WARN
	case zerolog.ErrorLevel:
		return log.ERROR
	case zerolog.FatalLevel:
		return 7 // log.fatalLevel
	case zerolog.PanicLevel:
		return 6 // log.panicLevel
	case zerolog.Disabled:
		return log.OFF
	default:
		return log.INFO
	}
}

func levelEchoToZero(l log.Lvl) zerolog.Level {
	switch l {
	case log.DEBUG:
		return zerolog.DebugLevel
	case log.WARN:
		return zerolog.WarnLevel
	case log.ERROR:
		return zerolog.ErrorLevel
	case 7: // log.fatalLevel
		return zerolog.FatalLevel
	case 6: // log.panicLevel
		return zerolog.PanicLevel
	case log.OFF:
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}
