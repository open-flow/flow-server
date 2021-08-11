package logger

import (
	"autoflow/pkg/common"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.SugaredLogger
}

func NewService(
	logger *zap.SugaredLogger,
) (*Service, error) {
	obj := &Service{
		logger,
	}

	obj.logger = obj.logger.With(zap.String("service", "logger"))

	return obj, nil
}

func (s *Service) Error(id common.ByProject, message string, data interface{}) {

}
