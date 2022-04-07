package zap

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

var _ log.Logger = (*Logger)(nil)

const (
	msgKey = "msg"
)

type Logger struct {
	log *zap.Logger
}

func NewLogger(zlog *zap.Logger) *Logger {
	return &Logger{zlog}
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	msg := ""
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		key := fmt.Sprint(keyvals[i])
		if key == msgKey {
			msg = fmt.Sprint(keyvals[i+1])
		} else {
			data = append(data, zap.Any(key, keyvals[i+1]))
		}
	}

	switch level {
	case log.LevelDebug:
		l.log.Debug(msg, data...)
	case log.LevelInfo:
		l.log.Info(msg, data...)
	case log.LevelWarn:
		l.log.Warn(msg, data...)
	case log.LevelError:
		l.log.Error(msg, data...)
	case log.LevelFatal:
		l.log.Fatal(msg, data...)
	}
	return nil
}

func (l *Logger) Sync() error {
	return l.log.Sync()
}
