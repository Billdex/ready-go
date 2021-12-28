package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var logger = NewLogger("json", "info")

func L() *zap.Logger {
	return logger
}

func Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Debugw(msg, keysAndValues...)
}

func Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Infow(msg, keysAndValues...)
}

func Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Warnw(msg, keysAndValues...)
}

func Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Errorw(msg, keysAndValues...)
}

func Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Sugar().Panic(args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Panicw(msg, keysAndValues...)
}

func Panicf(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Sugar().Fatal(args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Fatalw(msg, keysAndValues...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Sugar().Fatalf(template, args...)
}

func Sync() {
	logger.Sync()
}

func NewLogger(style string, level string, hooks ...func(zapcore.Entry) error) *zap.Logger {
	core := zapcore.NewCore(newEncoder(style), os.Stdout, getZapLevel(level))
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.Hooks(hooks...))
}

func SetLogger(style string, level string, hooks ...func(zapcore.Entry) error) {
	logger = NewLogger(style, level, hooks...)
}

// Zap 日志编码器
func newEncoder(style string) zapcore.Encoder {
	cfg := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "file",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	switch strings.ToLower(style) {
	case "console":
		return zapcore.NewConsoleEncoder(cfg)
	case "json":
		return zapcore.NewJSONEncoder(cfg)
	default:
		return zapcore.NewConsoleEncoder(cfg)
	}
}

// 转换 zap 日志级别
func getZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
