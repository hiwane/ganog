package ganog

import (
	"log"
	"runtime"
	"strings"
)

var logger Logger = NewLogger(
	log.New(log.Writer(), "", log.Ltime|log.Lmsgprefix),
	LevelInfo,
)

type BasicLogger struct {
	logger *log.Logger
	level  int
	format string
	color  bool
}

func NewLogger(logger *log.Logger, level int) *BasicLogger {
	return &BasicLogger{logger: logger, level: level, format: "%s:%4d: "}
}

func (l *BasicLogger) SetLevel(level int) {
	l.level = level
}

func (l *BasicLogger) Level() int {
	return l.level
}

/** %s, %d が必須 **/
func (l *BasicLogger) SetFormat(format string) {
	l.format = format
}

func (l *BasicLogger) Format() string {
	return l.format
}

func (l *BasicLogger) SetColor(color bool) {
	l.color = color
}

func (l *BasicLogger) get_color(level int) int {
	if !l.color {
		return -1
	}
	switch level {
	case LevelFatal:
		return 31
	case LevelError:
		return 31
	case LevelWarn:
		return 33
	case LevelInfo:
		return 36
	default:
		return -1
	}
}

func (l *BasicLogger) Log(level int, format string, v ...interface{}) {
	if l.level >= level {

		c := l.get_color(level)

		_, file, line, ok := runtime.Caller(2)

		anys := make([]any, 0, len(v)+5)
		if ok {
			// フルパスは長いので、ファイル名のみにする
			// 最後のスラッシュを探す
			if pos := strings.LastIndex(file, "/"); pos != -1 {
				file = file[pos+1:]
			}
			anys = append(anys, file, line)
			if c > 0 {
				anys = append(anys, c)
				format = l.format + "\x1b[%dm" + format + "\x1b[0m"
			} else {
				format = l.format + format
			}
		} else {
			if c > 0 {
				anys = append(anys, c)
				format = "\x1b[%dm" + format + "\x1b[0m"
			}
		}
		anys = append(anys, v...)
		l.logger.Printf(format, anys...)
	}
}

func (l *BasicLogger) Fatal(format string, v ...interface{}) {
	l.Log(LevelFatal, "FATAL:"+format, v...)
}

func (l *BasicLogger) Error(format string, v ...interface{}) {
	l.Log(LevelError, "ERROR:"+format, v...)
}

func (l *BasicLogger) Warn(format string, v ...interface{}) {
	l.Log(LevelWarn, "WARN :"+format, v...)
}

func (l *BasicLogger) Info(format string, v ...interface{}) {
	l.Log(LevelInfo, "INFO :"+format, v...)
}

func (l *BasicLogger) Debug(format string, v ...interface{}) {
	l.Log(LevelDebug, "DEBUG:"+format, v...)
}
