package logger

type Usecase interface {
	Info(format string, v ...any)
	Error(format string, v ...any)
}
