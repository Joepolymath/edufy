package common

import (
	"Learnium/internal/config"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

// describe logger behaviour
type ILogger interface {
	With(fields ...zap.Field) *Logger
	Error(ctx context.Context, msg string, fields ...zap.Field)
}

// Logger is a custom logger class and can be used as a dependency in controllers
type Logger struct {
	logger *zap.Logger
}

var cfg = config.Config

var logger *zap.Logger

type loggerfields string

var loggerFieldsAll loggerfields = "logger.fields"

// NewLogger initializes a new Logger instance
func NewLogger() *Logger {

	// Initialize Sentry with your DSN
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.SentryDSN,
		EnableTracing:    cfg.SentryEnableTracing,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return &Logger{logger: logger}
}

// With adds fields to the logger instance
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{logger: l.logger.With(fields...)}
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	defer l.ShutdownSentry()

	l.logWithContext(ctx, l.logger.Error, msg, fields...)
	// Capture the error and send it to Sentry
	sentry.CaptureException(errors.New(fmt.Sprintf("The Error: %s ", msg)))
}

// Info logs an info message
func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	defer l.ShutdownSentry()

	l.logWithContext(ctx, l.logger.Info, msg, fields...)
	sentry.CaptureMessage(fmt.Sprintf("The Info %s ", msg))
}

// logWithContext is a helper function to log messages with context and fields
func (l *Logger) logWithContext(ctx context.Context, logFunc func(msg string, fields ...zap.Field), msg string, fields ...zap.Field) {
	data := ctx.Value(loggerFieldsAll)
	var storedFields []zap.Field
	if data != nil {
		storedFields = data.([]zap.Field)
	}
	storedFields = append(storedFields, fields...)

	logFunc(msg, storedFields...)
}

// ShutdownSentry should be called when your application exits to flush Sentry events
func (l *Logger) ShutdownSentry() {
	sentry.Flush(2 * time.Second) // Adjust the timeout as needed
}

// Usage Example:
// logger := utils.NewLogger()
// logger.Info(context.TODO(), "Informational message", zap.String("key", "value"))
// logger.Error(context.TODO(), "Error message", zap.Error(err))
