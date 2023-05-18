package mock

import (
	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	return zap.Must(zap.NewProduction())
}
