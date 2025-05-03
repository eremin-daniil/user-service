package model

var (
	Stop              LogCommand = "Stop"
	RaiseLoggingLevel LogCommand = "RaiseLoggingLevel"
	LowerLoggingLevel LogCommand = "LowerLoggingLevel"
)

type LogCommand string
