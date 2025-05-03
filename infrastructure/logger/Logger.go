package logger

import (
	"context"
	loggerModel "user_service/infrastructure/logger/model"
)

type Logger struct {
	logChan chan<- *loggerModel.LogData
}

func NewLogger(logChan chan<- *loggerModel.LogData) *Logger {
	return &Logger{logChan: logChan}
}

func (l *Logger) Error(ctx context.Context, err error) {
	go l.sendData(loggerModel.ErrorLogData(ctx, err.Error()))
}

func (l *Logger) Errors(ctx context.Context, errs []error) {
	for _, err := range errs {
		go l.sendData(loggerModel.ErrorLogData(ctx, err.Error()))
	}
}

func (l *Logger) Warning(ctx context.Context, msg string) {
	go l.sendData(loggerModel.WarnLogData(ctx, msg))
}

func (l *Logger) Info(ctx context.Context, msg string) {
	go l.sendData(loggerModel.InfoLogData(ctx, msg))
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	go l.sendData(loggerModel.DebugLogData(ctx, msg))
}

func (l *Logger) sendData(logData *loggerModel.LogData) {
	l.logChan <- logData
}
