package file

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"user_service/infrastructure/constant"
	"user_service/infrastructure/logger"
	"user_service/infrastructure/logger/formatter"
	loggerModel "user_service/infrastructure/logger/model"
)

const LoggerID = "MainLogger"

const (
	defaultLevel              = loggerModel.InfoLevel
	defaultInputBufferSize    = 100
	defaultLoggerWriteTimeout = 100 * time.Millisecond
)

type MainLogger struct {
	inputLogChan   chan *loggerModel.LogData
	commandLogChan chan loggerModel.LogCommand
	level          loggerModel.LogLevel
	log            *logrus.Logger
}

func NewMainLogger(bufferSize int) *MainLogger {
	log := logrus.New()
	log.Formatter = formatter.NewTextFormatter()
	if bufferSize == 0 {
		bufferSize = defaultInputBufferSize
	}
	return &MainLogger{
		inputLogChan:   make(chan *loggerModel.LogData, bufferSize),
		commandLogChan: make(chan loggerModel.LogCommand),
		level:          defaultLevel,
		log:            log,
	}
}

func (ml *MainLogger) CommandLogChan() chan loggerModel.LogCommand {
	return ml.commandLogChan
}

func (ml *MainLogger) StartLogger() {
	ml.process(loggerModel.InfoLogData(context.Background(), fmt.Sprintf("%s запущен", LoggerID)))
	for {
		select {
		case logData := <-ml.inputLogChan:
			ml.process(logData)
		case commandLog := <-ml.commandLogChan:
			switch commandLog {
			case loggerModel.Stop:
				ml.log.Infof("Работа %s завершена командой: %s", LoggerID, loggerModel.Stop)
				return
			case loggerModel.LowerLoggingLevel:
				logLevel, err := ml.level.Lower()
				if err != nil {
					ml.log.Warnf("%s: %s", LoggerID, err)
					continue
				}
				ml.level = logLevel
			case loggerModel.RaiseLoggingLevel:
				logLevel, err := ml.level.Raise()
				if err != nil {
					ml.log.Warnf("%s: %s", LoggerID, err)
					continue
				}
				ml.level = logLevel
			}
		}
	}
}

func (ml *MainLogger) WriteLog(logData *loggerModel.LogData) {
	if ml.level > logData.Level() {
		return
	}
	select {
	case ml.inputLogChan <- logData:
	case <-time.After(defaultLoggerWriteTimeout):
		ml.log.Errorf("Превышено время записи в буфер для %s", LoggerID)
	}
}

func (ml *MainLogger) process(logData *loggerModel.LogData) {
	entry := logrus.NewEntry(ml.log)
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
