package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	GetDefer() func()
}

type MyLogger struct {
	log zap.Logger
}

func (l *MyLogger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *MyLogger) Info(msg string) {
	l.log.Info(msg)
}

func (l *MyLogger) Error(msg string) {
	l.log.Error(msg)
}

func (l *MyLogger) GetDefer() func() {
	return func() {
		l.log.Sync()
	}
}

func New(debug bool) (Logger, error) {
	var logger *zap.Logger
	var err error
	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}
	return &MyLogger{
		log: *logger,
	}, nil
}
