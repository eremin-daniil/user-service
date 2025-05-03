package model

import "context"

type LogData struct {
	ctx   context.Context
	msg   string
	level LogLevel
}

func ErrorLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		ctx:   ctx,
		msg:   msg,
		level: ErrorLevel,
	}
}

func WarnLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		ctx:   ctx,
		msg:   msg,
		level: WarnLevel,
	}
}

func InfoLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		ctx:   ctx,
		msg:   msg,
		level: InfoLevel,
	}
}

func DebugLogData(ctx context.Context, msg string) *LogData {
	return &LogData{
		ctx:   ctx,
		msg:   msg,
		level: DebugLevel,
	}
}

func (l *LogData) Ctx() context.Context {
	return l.ctx
}

func (l *LogData) Msg() string {
	return l.msg
}

func (l *LogData) Level() LogLevel {
	return l.level
}
