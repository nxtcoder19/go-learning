package main

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger interface {
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Errorf(err error, args ...interface{})
	Warnf(msg string, args ...interface{})
	WithName(name string) Logger
	WithKV(keyValuePairs ...interface{}) Logger
}

type logger struct {
	zapLogger *zap.SugaredLogger
}

func NewLogger() Logger {
	zapLogger, _ := zap.NewDevelopment()
	return &logger{
		zapLogger: zapLogger.Sugar(),
	}
}

func (l *logger) Debugf(msg string, args ...interface{}) {
	l.zapLogger.Debugf(msg, args...)
}

func (l *logger) Infof(msg string, args ...interface{}) {
	l.zapLogger.Infof(msg, args...)
}

func (l *logger) Errorf(err error, args ...interface{}) {
	args = append(args, zap.Error(err))
	l.zapLogger.Errorf(msgWithErr(err, args...))
}

func (l *logger) Warnf(msg string, args ...interface{}) {
	l.zapLogger.Warnf(msg, args...)
}

func (l *logger) WithName(name string) Logger {
	return &logger{
		zapLogger: l.zapLogger.Named(name),
	}
}

func (l *logger) WithKV(keyValuePairs ...interface{}) Logger {
	return &logger{
		zapLogger: l.zapLogger.With(keyValuePairs...),
	}
}

func msgWithErr(err error, args ...interface{}) string {
	return fmt.Sprintf("%s - %v", fmt.Sprint(args...), err)
}

func main() {
	myLogger := NewLogger()

	myLogger.Debugf("Debug message: %s", "debug info")
	myLogger.Infof("Info message: %s", "info info")

	err := fmt.Errorf("an error occurred")
	myLogger.Errorf(err, "Error message: %s", "error info")

	myLogger.WithName("Example").Warnf("Warning message: %s", "warning info")
	myLogger.WithKV("key", "value").Infof("Log with key-value pair")
}
