package logger

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{}
}

func (l *LogrusLogger) Info(msg ...interface{}) {
	logrus.Info(msg...)

	go l.writeToFile()
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)

	go l.writeToFile()
}

func (l *LogrusLogger) Warn(msg ...interface{}) {
	logrus.Warn(msg...)

	go l.writeToFile()
}

func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)

	go l.writeToFile()
}

func (l *LogrusLogger) Error(msg ...interface{}) {
	logrus.Error(msg...)

	go l.writeToFile()
}

func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)

	go l.writeToFile()
}

func (l *LogrusLogger) Fatal(msg ...interface{}) {
	logrus.Fatal(msg...)

	go l.writeToFile()
}

func (l *LogrusLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)

	go l.writeToFile()
}

func (l *LogrusLogger) writeToFile() {
	// TODO write to file feature
}
