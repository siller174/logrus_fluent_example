package main

import (
	"github.com/evalphobia/logrus_fluent"
	"github.com/sirupsen/logrus"
)

func main() {
	Debug("This debug message")
	Info("This Info message")
	Warn("This warn message")
	Error("This error message")
	Fatal("This fatal message")
}

func init() {
	hook, err := logrus_fluent.NewWithConfig(logrus_fluent.Config{
		Host: "localhost",
		Port: 9880,
	})
	if err != nil {
		panic(err)
	}
	hook.SetTag("fluent_example")
	hook.SetMessageField("message")
	logrus.AddHook(hook)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func Debug(message string, args ...interface{}) {
	logrus.Debugf(message, args...)
}

func Info(message string, args ...interface{}) {
	logrus.Infof(message, args...)
}

func Warn(message string, args ...interface{}) {
	logrus.Warnf(message, args...)
}

func Error(message string, args ...interface{}) {
	logrus.Errorf(message, args...)
}

func Fatal(message string, args ...interface{}) {
	logrus.Fatalf(message, args...)
}
