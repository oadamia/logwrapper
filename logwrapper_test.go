package logwrapper

import (
	"os"
	"testing"
	"time"

	"github.com/oadamia/test"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type testWriter struct {
	output string
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.output = string(p[:])
	return len(p), nil
}

func TestMain(m *testing.M) {
	Init(config())
	setTimestampFunc(mock_timestampFunc)
	setTimestampFieldName("@timestamp")
	setCallerMarshalFunction(mock_callerMarshalFunc)

	os.Exit(m.Run())
}

func config() Config {
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
	file = "/Users/oto/Projects/microsena/logwrapper/loggerwrapper_test.go"
	return callerMarshalFunc(pc, file, 73)
}

func TestLoggerPackage(t *testing.T) {

	t.Run("Config", func(t *testing.T) {
		assert := assert.New(t)
		c := config()
		wrapper = Wrapper()
		assert.Equal(c.Level, logger.GetLevel().String())
		assert.Equal(len(writers), 2)
		assert.Equal(c.TimeFieldFormat, zerolog.TimeFieldFormat)
		assert.Equal("tester.log", fileName(c.FileName, c.FilePath))
	})
}

func TestWrapper(t *testing.T) {

	tw := new(testWriter)

	t.Run("Output", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.SetOutput(tw)
		assert.Equal(multiWriter, wrapper.Output())
	})

	t.Run("Prefix", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.SetPrefix("Prefix")
		assert.Equal("Prefix", wrapper.Prefix())
		// TODO Check if Prefix is written
	})

	t.Run("Print", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Print("test")
		assert.Equal(test.Read("json/Print.json"), tw.output)
	})

	t.Run("Printf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Printf("t:%s", "test")
		assert.Equal(test.Read("json/Printf.json"), tw.output)
	})

	t.Run("Debug", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Debug("test")
		assert.Equal(test.Read("json/wrapperDebug.json"), tw.output)
	})

	t.Run("Debugf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Debugf("t:%s", "test")
		assert.Equal(test.Read("json/wrapperDebugf.json"), tw.output)
	})

	t.Run("Info", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Info("test")
		assert.Equal(test.Read("json/wrapperInfo.json"), tw.output)
	})

	t.Run("Infof", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Infof("t:%s", "test")
		assert.Equal(test.Read("json/wrapperInfof.json"), tw.output)
	})

	t.Run("Warn", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Warn("test")
		assert.Equal(test.Read("json/wrapperWarn.json"), tw.output)
	})

	t.Run("Warnf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Warnf("t:%s", "test")
		assert.Equal(test.Read("json/wrapperWarnf.json"), tw.output)
	})

	t.Run("Err", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Err(nil)
		//because writer is not overwriten
		assert.Equal(test.Read("json/wrapperWarnf.json"), tw.output)
	})

	t.Run("Error", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Error("test")
		assert.Equal(test.Read("json/wrapperError.json"), tw.output)
	})

	t.Run("Errorf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Errorf("t:%s", "test")
		assert.Equal(test.Read("json/wrapperErrorf.json"), tw.output)
	})

}
