package log

import (
	"fmt"
	"io"
	golog "log"
	"os"
)

type WOFLogger struct {
	Logger   *golog.Logger
	MinLevel string
	levels   map[string]int
}

func NewWOFLogger(out io.Writer, prefix string, minlevel string) *WOFLogger {

	levels := make(map[string]int)
	levels["fatal"] = 0
	levels["error"] = 10
	levels["warning"] = 20
	levels["status"] = 25
	levels["info"] = 30
	levels["debug"] = 40

	logger := golog.New(out, prefix, golog.Lmicroseconds)

	l := WOFLogger{Logger: logger, MinLevel: minlevel, levels: levels}
	return &l
}

func (l WOFLogger) Debug(format string, v ...interface{}) {
	l.format("debug", format, v...)
}

func (l WOFLogger) Info(format string, v ...interface{}) {
	l.format("info", format, v...)
}

func (l WOFLogger) Status(format string, v ...interface{}) {
	l.format("status", format, v...)
}

func (l WOFLogger) Warning(format string, v ...interface{}) {
	l.format("warning", format, v...)
}

func (l WOFLogger) Error(format string, v ...interface{}) {
	l.format("error", format, v...)
}

func (l WOFLogger) Fatal(format string, v ...interface{}) {
	l.format("fatal", format, v...)
	os.Exit(1)
}

func (l WOFLogger) format(level string, format string, v ...interface{}) {

	if l.emit(level) {

		msg := fmt.Sprintf(format, v...)
		out := fmt.Sprintf("[%s] %s", level, msg)
		l.Logger.Println(out)
	}
}

func (l WOFLogger) emit(level string) bool {

	this_level, ok := l.levels[level]

	if !ok {
		return false
	}

	min_level, ok := l.levels[l.MinLevel]

	if !ok {
		return false
	}

	if this_level > min_level {
		return false
	}

	return true
}
