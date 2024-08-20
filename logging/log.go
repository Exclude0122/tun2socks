package logging

import (
	"go.uber.org/zap"
)

func init() {
	logger := zap.NewProductionConfig()
	l := logger.Build()
}
