package file

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"user_service/infrastructure/constant"
	"user_service/infrastructure/logger"
	loggerModel "user_service/infrastructure/logger/model"
)

const LoggerID = "FileLogger"

const (
	defaultLevel              = loggerModel.InfoLevel
	defaultInputBufferSize    = 100
	defaultLoggerWriteTimeout = 100 * time.Millisecond
)

type FileLogger struct {
	appID          string
	inputLogChan   chan *loggerModel.LogData
	commandLogChan chan loggerModel.LogCommand
	level          loggerModel.LogLevel
	log            *logrus.Logger
}

func NewFileLogger(file *os.File, appID string, bufferSize int) *FileLogger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)
	if bufferSize == 0 {
		bufferSize = defaultInputBufferSize
	}
	return &FileLogger{
		appID:          appID,
		inputLogChan:   make(chan *loggerModel.LogData, bufferSize),
		commandLogChan: make(chan loggerModel.LogCommand),
		level:          defaultLevel,
		log:            log,
	}
}

func (fl *FileLogger) CommandLogChan() chan loggerModel.LogCommand {
	return fl.commandLogChan
}

func (fl *FileLogger) StartLogger() {
	fl.process(loggerModel.InfoLogData(context.Background(), fmt.Sprintf("%s запущен", LoggerID)))
	for {
		select {
		case logData := <-fl.inputLogChan:
			fl.process(logData)
		case commandLog := <-fl.commandLogChan:
			switch commandLog {
			case loggerModel.Stop:
				fl.log.Infof("Работа %s завершена командой: %s", LoggerID, loggerModel.Stop)
				return
			case loggerModel.LowerLoggingLevel:
				logLevel, err := fl.level.Lower()
				if err != nil {
					fl.log.Warnf("%s: %s", LoggerID, err)
					continue
				}
				fl.level = logLevel
			case loggerModel.RaiseLoggingLevel:
				logLevel, err := fl.level.Raise()
				if err != nil {
					fl.log.Warnf("%s: %s", LoggerID, err)
					continue
				}
				fl.level = logLevel
			}
		}
	}
}

func (fl *FileLogger) WriteLog(logData *loggerModel.LogData) {
	if fl.level > logData.Level() {
		return
	}
	select {
	case fl.inputLogChan <- logData:
	case <-time.After(defaultLoggerWriteTimeout):
		fl.log.Errorf("Превышено время записи в буфер для %s", LoggerID)
	}
}

func (fl *FileLogger) process(logData *loggerModel.LogData) {
	entry := fl.log.WithField(logger.AppIDFieldKey, fl.appID)
	requestID := logData.Ctx().Value(constant.RequestIDCtxKey)
	if requestID != nil {
		entry = entry.WithField(logger.RequestIDFieldKey, requestID)
	}

	switch logData.Level() {
	case loggerModel.DebugLevel:
		entry.Debug(logData.Msg())
	case loggerModel.InfoLevel:
		entry.Info(logData.Msg())
	case loggerModel.WarnLevel:
		entry.Warn(logData.Msg())
	case loggerModel.ErrorLevel:
		entry.Error(logData.Msg())
	}
}
