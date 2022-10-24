package logwrapper

import (
	"errors"
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
	setOpenFileFunction(mock_openFileFunc)

	Init(mock_config())
	setTimestampFunc(mock_timestampFunc)
	setTimestampFieldName("@timestamp")
	setCallerMarshalFunction(mock_callerMarshalFunc)

	os.Exit(m.Run())
}

func TestWrapper(t *testing.T) {

	tw := new(testWriter)

	t.Run("TimestampFunc", func(t *testing.T) {
		assert := assert.New(t)
		assert.GreaterOrEqual(time.Now().UTC(), utcTimeFunc())
	})

	t.Run("Config", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal("trace", logger.GetLevel().String())
		assert.Equal(2, len(writers))
		assert.Equal("2006-01-02T15:04:05.999999", zerolog.TimeFieldFormat)
		assert.Equal("test/tester.log", fileName("tester.log", "test"))
	})

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
		assert.Equal(test.Read("testdata/Print.json"), tw.output)
	})

	t.Run("Printf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Printf("t:%s", "test")
		assert.Equal(test.Read("testdata/Printf.json"), tw.output)
	})

	t.Run("Debug", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Debug("test")
		assert.Equal(test.Read("testdata/wrapperDebug.json"), tw.output)
	})

	t.Run("Debugf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Debugf("t:%s", "test")
		assert.Equal(test.Read("testdata/wrapperDebugf.json"), tw.output)
	})

	t.Run("Info", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Info("test")
		assert.Equal(test.Read("testdata/wrapperInfo.json"), tw.output)
	})

	t.Run("Infof", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Infof("t:%s", "test")
		assert.Equal(test.Read("testdata/wrapperInfof.json"), tw.output)
	})

	t.Run("Warn", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Warn("test")
		assert.Equal(test.Read("testdata/wrapperWarn.json"), tw.output)
	})

	t.Run("Warnf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Warnf("t:%s", "test")
		assert.Equal(test.Read("testdata/wrapperWarnf.json"), tw.output)
	})

	t.Run("Err", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Err(nil)
		//because writer is not overwriten
		assert.Equal(test.Read("testdata/wrapperWarnf.json"), tw.output)
	})

	t.Run("Error", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Error("test")
		assert.Equal(test.Read("testdata/wrapperError.json"), tw.output)
	})

	t.Run("Errorf", func(t *testing.T) {
		assert := assert.New(t)

		wrapper.Errorf("t:%s", "test")
		assert.Equal(test.Read("testdata/wrapperErrorf.json"), tw.output)
	})

	t.Run("Fatal", func(t *testing.T) {
		assert := assert.New(t)
		loggerFatal = mock_loggerFatal
		defer func() {
			loggerFatal = logger.Fatal
		}()
		wrapper.Fatal("fatal")
		assert.Equal(test.Read("testdata/wrapperFatal.json"), tw.output)
	})

	t.Run("Fatalf", func(t *testing.T) {
		assert := assert.New(t)
		loggerFatal = mock_loggerFatal
		defer func() {
			loggerFatal = logger.Fatal
		}()
		wrapper.Fatalf("t:%s", "fatal")
		assert.Equal(test.Read("testdata/wrapperFatalf.json"), tw.output)
	})

	t.Run("Panic", func(t *testing.T) {
		assert := assert.New(t)
		loggerPanic = mock_loggerPanic
		defer func() {
			loggerPanic = logger.Panic
		}()
		wrapper.Panic("panic")
		assert.Equal(test.Read("testdata/wrapperPanic.json"), tw.output)
	})

	t.Run("Panicf", func(t *testing.T) {
		assert := assert.New(t)
		loggerPanic = mock_loggerPanic
		defer func() {
			loggerPanic = logger.Panic
		}()
		wrapper.Panicf("t:%s", "panic")
		assert.Equal(test.Read("testdata/wrapperPanicf.json"), tw.output)
	})

}
