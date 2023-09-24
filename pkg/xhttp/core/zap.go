package core

//go:generate mockgen -destination=mocks/mock_logger.go -package=mocks github.com/jedi-knights/xhttp/pkg/xhttp/core LoggerInterface

import (
	"go.uber.org/zap/zapcore"
)

type LoggerInterface interface {
	Info(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
}
