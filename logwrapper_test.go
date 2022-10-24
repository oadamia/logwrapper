package logwrapper

import (
	"errors"
	"os"
	"testing"

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
	Init(mock_config())
	setTimestampFunc(mock_timestampFunc)
	setTimestampFieldName("@timestamp")
	setCallerMarshalFunction(mock_callerMarshalFunc)
	setOpenFileFunction(mock_openFileFunc)

	os.Exit(m.Run())
}

func TestWrapper(t *testing.T) {

	tw := new(testWriter)

	t.Run("Config open file Error", func(t *testing.T) {
		assert := assert.New(t)
		config := mock_config()
		config.FileName = "error"
		err := setFileOutput(config)
		if assert.Error(err) {
			assert.EqualError(err, "open file error")
		}
	})

	t.Run("Config file stat Error", func(t *testing.T) {
		assert := assert.New(t)
		config := mock_config()
		config.FileName = "Stat error"
		err := setFileOutput(config)
		if assert.Error(err) {
			assert.EqualError(err, "invalid argument")
		}
	})

	t.Run("Config", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal("trace", logger.GetLevel().String())
		assert.Equal(2, len(writers))
		assert.Equal("2006-01-02T15:04:05.999999", zerolog.TimeFieldFormat)
		assert.Equal("test/tester.log", fileName("tester.log", "test"))
	})

	t.Run("wrapper", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(wrapper, Wrapper())

	})

	t.Run("Output", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.SetOutput(tw)
		assert.Equal(multiWriter, wrapper.Output())
	})

	t.Run("Prefix", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.SetPrefix("Prefix")
		assert.Equal("Prefix", wrapper.Prefix())

	})

	t.Run("Header", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.SetHeader("Header")
		assert.Equal("{\"level\":\"error\",\"@timestamp\":\"2008-01-08T17:05:05\",\"message\":\"SetHeader is not implemented\"}\n", tw.output)
	})

	t.Run("err", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Err(errors.New("error"))
		assert.Equal("{\"level\":\"error\",\"caller\":\"Projects/microsena/logwrapper/loggerwrapper_test.go:73\",\"@timestamp\":\"2008-01-08T17:05:05\",\"message\":\"error\"}\n", tw.output)
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
