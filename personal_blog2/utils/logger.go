package utils

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger 初始化日志配置
func InitLogger() {
	// 配置日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	logLevel := zapcore.InfoLevel
	if logLevelStr := os.Getenv("LOG_LEVEL"); logLevelStr != "" {
		parsedLevel, err := zapcore.ParseLevel(logLevelStr)
		if err != nil {
			// 如果解析失败，记录错误并使用默认级别
			zap.L().Error("Failed to parse log level, using default info level", zap.Error(err))
		} else {
			logLevel = parsedLevel
		}
	}

	// 创建Core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 输出到控制台
		logLevel,
	)

	// 创建Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// 替换zap全局Logger
	zap.ReplaceGlobals(Logger)
}

// SyncLogger 刷新日志缓冲区
func SyncLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}
