package logwrapper

import (
	"github.com/rs/zerolog"
)

// package wide logger variable, every log operation should use logger instance
var logger zerolog.Logger

// wrapper type
type logWrapper struct {
	prefix string
}
