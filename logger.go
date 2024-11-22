package ganog

const (
	LevelNone  = 0
	LevelFatal = 1
	LevelError = 2
	LevelWarn  = 3
	LevelInfo  = 4
	LevelDebug = 5
)

type Logger interface {
	SetLevel(level int)
	Log(level int, format string, v ...interface{})

	Fatal(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

func Log(level int, format string, v ...interface{}) {
	logger.Log(level, format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Fatal(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Warn(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Debug(format, v...)
}

func SetLogger(l Logger) {
	logger = l
}
