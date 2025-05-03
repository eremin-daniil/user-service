package logger

import (
	loggerInterface "user_service/infrastructure/logger/interface"
	loggerModel "user_service/infrastructure/logger/model"
)

const defaultInputBufferSize = 100

type Service struct {
	inputChan chan *loggerModel.LogData
	loggers   map[string]loggerInterface.Publisher
}

func NewService(bufferSize int) *Service {
	if bufferSize == 0 {
		bufferSize = defaultInputBufferSize
	}
	return &Service{
		inputChan: make(chan *loggerModel.LogData, bufferSize),
		loggers:   make(map[string]loggerInterface.Publisher),
	}
}

func (s *Service) InputChan() chan *loggerModel.LogData {
	return s.inputChan
}

func (s *Service) RegisterLogger(loggerID string, logger loggerInterface.Publisher) {
	s.loggers[loggerID] = logger
}

func (s *Service) Start() {
	go s.runMainWorker()

	for _, logger := range s.loggers {
		go logger.StartLogger()
	}
}

func (s *Service) runMainWorker() {
	for {
		select {
		case logData := <-s.inputChan:
			if logData == nil {
				continue
			}
			for _, logger := range s.loggers {
				logger.WriteLog(logData)
			}
		}
	}
}
