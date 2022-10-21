package logwrapper

var wrapper *logWrapper

// Init should be called before any usuge of logwrapper or behaviour will be undefined
func Init(c Config) error {
	wrapper = new(logWrapper)
	return configure(c)
}

// Wrapper get logger wrapper
func Wrapper() *logWrapper {
	return wrapper
}
