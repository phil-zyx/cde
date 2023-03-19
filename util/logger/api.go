package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _atomicLevel zap.AtomicLevel
var _appLogger *zap.SugaredLogger

func init() {
	_atomicLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	cfg := zap.Config{
		Level:             _atomicLevel,
		Development:       true,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{stdout},
		DisableCaller:     true,
		DisableStacktrace: true,
	}
	l, _ := cfg.Build()
	_appLogger = l.Sugar()
}

// Init 初始化日志模块
func Init(logFile string) error {
	var err error
	_appLogger, err = newAppLogger(logFile)
	if err != nil {
		return fmt.Errorf("new app logger error: %v", err)
	}
	return nil
}

// SetLogLevel 设置日志等级 debug|info|error
func SetLogLevel(level string) error {
	var l zapcore.Level
	err := l.UnmarshalText([]byte(level))
	if err != nil {
		return err
	}
	_atomicLevel.SetLevel(l)
	return err
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	_appLogger.Debug(args...)
}

// Debugf .
func Debugf(template string, args ...interface{}) {
	_appLogger.Debugf(template, args)
}

func Info(args ...interface{}) {
	_appLogger.Info(args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	_appLogger.Fatalf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	_appLogger.Errorf(template, args...)
}
