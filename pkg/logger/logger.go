package logger

type Logger interface {
	Info(msg ...interface{})
	Infof(format string, args ...interface{})
	Warn(msg ...interface{})
	Warnf(format string, args ...interface{})
	Error(msg ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(msg ...interface{})
	Fatalf(format string, args ...interface{})
}
