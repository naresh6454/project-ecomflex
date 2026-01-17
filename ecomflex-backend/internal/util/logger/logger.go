package logger

import (
	"fmt"
	"os"
	"strings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger interface defines the logging methods
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, err error, args ...interface{})
	Fatal(msg string, err error, args ...interface{})
}

// zapLogger implements Logger interface using zap
type zapLogger struct {
	logger *zap.Logger
}

// NewLogger creates a new logger
func NewLogger(level string) Logger {
	// Parse log level
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	case "fatal":
		zapLevel = zapcore.FatalLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// Create encoder config
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.CallerKey = "caller"
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	// Create core
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapLevel,
	)

	// Create logger
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &zapLogger{
		logger: logger,
	}
}

// Debug logs a debug message
func (l *zapLogger) Debug(msg string, args ...interface{}) {
	fields := argsToFields(args...)
	if len(args) > 0 && len(fields) == 0 {
		l.logger.Debug(fmt.Sprintf(msg, args...))
	} else {
		l.logger.Debug(msg, fields...)
	}
}

// Info logs an info message
func (l *zapLogger) Info(msg string, args ...interface{}) {
	fields := argsToFields(args...)
	if len(args) > 0 && len(fields) == 0 {
		l.logger.Info(fmt.Sprintf(msg, args...))
	} else {
		l.logger.Info(msg, fields...)
	}
}

// Warn logs a warning message
func (l *zapLogger) Warn(msg string, args ...interface{}) {
	fields := argsToFields(args...)
	if len(args) > 0 && len(fields) == 0 {
		l.logger.Warn(fmt.Sprintf(msg, args...))
	} else {
		l.logger.Warn(msg, fields...)
	}
}

// Error logs an error message
func (l *zapLogger) Error(msg string, err error, args ...interface{}) {
	fields := argsToFields(args...)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	
	var expectedFieldCount int
	if err != nil {
		expectedFieldCount = 1
	} else {
		expectedFieldCount = 0
	}
	
	if len(args) > 0 && len(fields) == expectedFieldCount {
		l.logger.Error(fmt.Sprintf(msg, args...))
	} else {
		l.logger.Error(msg, fields...)
	}
}

// Fatal logs a fatal message and exits
func (l *zapLogger) Fatal(msg string, err error, args ...interface{}) {
	fields := argsToFields(args...)
	if err != nil {
		fields = append(fields, zap.Error(err))
	}
	
	var expectedFieldCount int
	if err != nil {
		expectedFieldCount = 1
	} else {
		expectedFieldCount = 0
	}
	
	if len(args) > 0 && len(fields) == expectedFieldCount {
		l.logger.Fatal(fmt.Sprintf(msg, args...))
	} else {
		l.logger.Fatal(msg, fields...)
	}
}

// argsToFields converts args to zap fields
func argsToFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args)/2)
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key, ok := args[i].(string)
			if !ok {
				continue
			}
			fields = append(fields, zap.Any(key, args[i+1]))
		}
	}
	return fields
}