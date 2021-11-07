package loggers

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var config zap.Config

// cloudLoggingLevelEncoder logger level encoder for Google Cloud Logging
func cloudLoggingLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var errString string
	switch l {
	case zapcore.DebugLevel:
		errString = "DEBUG"
	case zapcore.InfoLevel:
		errString = "INFO"
	case zapcore.WarnLevel:
		errString = "WARNING"
	case zapcore.ErrorLevel:
		errString = "ERROR"
	case zapcore.DPanicLevel:
		errString = "ALERT"
	case zapcore.PanicLevel:
		errString = "CRITICAL"
	case zapcore.FatalLevel:
		errString = "EMERGENCY"
	default:
		errString = "DEFAULT"
	}
	enc.AppendString(errString)
}

// BuildLogger ...
func BuildLogger(logLevel zapcore.Level) *zap.Logger {
	config = zap.NewProductionConfig()
	config.Level.SetLevel(logLevel)

	config.EncoderConfig.LevelKey = "severity"
	config.EncoderConfig.EncodeLevel = cloudLoggingLevelEncoder

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return logger
}

// GetLoggerLevel ...
func GetLoggerLevel() string {
	return config.Level.Level().CapitalString()
}

// SetLoggerLevel ...
func SetLoggerLevel(level zapcore.Level) {
	config.Level.SetLevel(level)
}
