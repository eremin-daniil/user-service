package _interface

import loggerModel "user_service/infrastructure/logger/model"

type Publisher interface {
	StartLogger()
	WriteLog(logData *loggerModel.LogData)
	CommandLogChan() chan loggerModel.LogCommand
}
