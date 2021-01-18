package utils

import (
	"fmt"
	"time"
)

type Logger struct {
	Prefix        string
	ShowTime      bool
	ShowPrefix    bool
	ShowLevel     bool
	DEBUG_ENABLED bool
	INFO_ENABLED  bool
	WARN_ENABLED  bool
	ERROR_ENABLED bool
	TRACE_ENABLED bool
}

func NewLogger(prefix string) *Logger {
	l := Logger{Prefix: prefix}
	l.ShowTime = true
	l.ShowPrefix = true
	l.ShowLevel = true
	l.TRACE_ENABLED = false
	l.DEBUG_ENABLED = true
	l.WARN_ENABLED = true
	l.ERROR_ENABLED = true
	l.INFO_ENABLED = true
	return &l
}

func (l *Logger) StdOut(msg string) {
	fmt.Printf("%v\n", msg)
}

func (l *Logger) Trace(msg string) {
	if l.TRACE_ENABLED {
		msg2 := l.Decorate("Trace", msg)
		l.StdOut(msg2)
	}
}

func (l *Logger) Debug(msg string) {
	if l.DEBUG_ENABLED {
		msg2 := l.Decorate("Debug", msg)
		l.StdOut(msg2)
	}
}

func (l *Logger) Info(msg string) {
	if l.INFO_ENABLED {
		msg2 := l.Decorate("Info", msg)
		l.StdOut(msg2)
	}
}

func (l *Logger) Warn(msg string) {
	if l.WARN_ENABLED {
		msg2 := l.Decorate("Warn", msg)
		l.StdOut(msg2)
	}
}

func (l *Logger) Error(msg string) {
	if l.ERROR_ENABLED {
		msg2 := l.Decorate("Error", msg)
		l.StdOut(msg2)
	}
}

func (l *Logger) Decorate(level string, msg string) string {
	if l.ShowPrefix {
		msg = fmt.Sprintf("%v %v", l.Prefix, msg)
	}
	if l.ShowLevel {
		msg = fmt.Sprintf("%v %v", level, msg)
	}
	if l.ShowTime {
		t := time.Now()
		msg = fmt.Sprintf("%v %v", t.Format(time.RFC3339), msg)
	}
	return msg

}
