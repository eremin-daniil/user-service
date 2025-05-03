package model

import (
	"errors"
	"fmt"
)

type LogLevel int8
type LogLevelEnum map[LogLevel]string

const (
	DebugLevel LogLevel = 1
	InfoLevel  LogLevel = 2
	WarnLevel  LogLevel = 3
	ErrorLevel LogLevel = 4
)

var Levels = LogLevelEnum{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
}

func (l LogLevel) Lower() (LogLevel, error) {
	if l == DebugLevel {
		return l, errors.New(fmt.Sprintf("Невозможно опустить уровень логирования ниже %s", Levels[DebugLevel]))
	}
	return l - 1, nil
}

func (l LogLevel) Raise() (LogLevel, error) {
	if l == ErrorLevel {
		return l, errors.New(fmt.Sprintf("Невозможно поднять уровень логирования выше %s", Levels[ErrorLevel]))
	}
	return l + 1, nil
}
