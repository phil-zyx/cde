package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"net/url"
)

const (
	udpSink        = "udp"
	rotatefileSink = "rfile"
	stdout         = "stdout"
	stderr         = "stderr"
)

// newAppLogger 创建应用日志Logger
func newAppLogger(logFile string) (*zap.SugaredLogger, error) {
	err := SetLogLevel(zap.DebugLevel.String())
	if err != nil {
		return nil, err
	}
	err = regSinkAndEncoder(logFile)
	if err != nil {
		return nil, err
	}
	cfg := zap.Config{
		Level:             _atomicLevel,
		Development:       true,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{logFile},
		ErrorOutputPaths:  []string{stderr},
		DisableCaller:     false,
		DisableStacktrace: true,
	}
	cfg.EncoderConfig.EncodeTime = timeEncoder
	cfg.EncoderConfig.EncodeLevel = levelEncoder

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

// regSinkAndEncoder 注册相应的URL解析器和编码器
func regSinkAndEncoder(logFile string) error {
	err := zap.RegisterSink(rotatefileSink, func(_ *url.URL) (zap.Sink, error) {
		w := &rfileWriter{
			Logger: &lumberjack.Logger{
				Filename:   logFile[len(rotatefileSink)+1:],
				MaxSize:    500, // megabytes
				MaxBackups: 3,
				MaxAge:     28, // days
			}}
		return w, nil
	})
	if err != nil {
		return err
	}
	return nil
}
