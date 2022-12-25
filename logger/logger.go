package logger

var AppLog Logger

type Logger interface {
	Debug(ctx ...interface{})
	Info(ctx ...interface{})
	Warn(ctx ...interface{})
	Error(ctx ...interface{})
}

func SetAppLog(logger Logger) {
	AppLog = logger
}
