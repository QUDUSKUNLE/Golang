package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// InitLogger initializes the global logger
func InitLogger() {
	var err error
	// Use zap's production configuration
	log, err = zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

// Sync flushes any buffered log entries
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
