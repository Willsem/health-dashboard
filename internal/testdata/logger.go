package testdata

import (
	"github.com/Willsem/health-dashboard/internal/logger"
	"github.com/Willsem/health-dashboard/internal/logger/zap"
)

type testWriter struct{}

func (testWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func NewLogger() logger.Logger {
	return zap.NewLogger(zap.Output(testWriter{}))
}
