package _interface

import "context"

type Logger interface {
	Error(ctx context.Context, err error)
	Errors(ctx context.Context, errs []error)
	Warning(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Debug(ctx context.Context, message string)
}
