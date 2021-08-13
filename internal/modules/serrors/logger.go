package serrors

import (
	"autoflow/pkg/common"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger(
	logger *zap.SugaredLogger,
) (*Logger, error) {
	obj := &Logger{
		logger,
	}

	obj.logger = obj.logger.With(zap.String("service", "logger"))

	return obj, nil
}

func (s *Logger) Error(id common.SpacedObject, message string, data interface{}) {

}
