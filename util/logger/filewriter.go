package logger


import "github.com/natefinch/lumberjack"

// rotate file writer
type rfileWriter struct {
	*lumberjack.Logger
}

// Sync implements zap.Sink
func (rfileWriter) Sync() error {
	return nil
}