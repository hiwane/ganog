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

func (l *BasicLogger) Log(level int, format string, v ...interface{}) {
	if l.level >= level {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			// フルパスは長いので、ファイル名のみにする
			// 最後のスラッシュを探す
			if pos := strings.LastIndex(file, "/"); pos != -1 {
				file = file[pos+1:]
			}
			anys := make([]interface{}, 0, len(v)+2)
			anys = append(anys, file, line)
			anys = append(anys, v...)
			l.logger.Printf(l.format+format, anys...)
		} else {
			l.logger.Printf(format, v...)
		}
	}
}

func (l *BasicLogger) Fatal(format string, v ...interface{}) {
	l.Log(LevelFatal, format, v...)
}

func (l *BasicLogger) Error(format string, v ...interface{}) {
	l.Log(LevelError, format, v...)
}

func (l *BasicLogger) Warn(format string, v ...interface{}) {
	l.Log(LevelWarn, format, v...)
}

func (l *BasicLogger) Info(format string, v ...interface{}) {
	l.Log(LevelInfo, format, v...)
}

func (l *BasicLogger) Debug(format string, v ...interface{}) {
	l.Log(LevelDebug, format, v...)
}
