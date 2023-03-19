package logger

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"time"
)

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(formatTime(t))
}

func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.000000")
}

func levelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(formatLevel(level))
}

func formatLevel(level zapcore.Level) string {
	var s string
	switch level {
	case zapcore.DebugLevel:
		s = "debug"
	case zapcore.InfoLevel:
		s = "info"
	case zapcore.WarnLevel:
		s = "warning"
	case zapcore.ErrorLevel:
		s = "error"
	case zapcore.DPanicLevel, zapcore.PanicLevel:
		s = "critical"
	case zapcore.FatalLevel:
		s = "emerg"
	default:
		s = fmt.Sprintf("Level(%d)", level)
	}
	return s
}