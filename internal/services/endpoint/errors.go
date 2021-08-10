package endpoint

import (
	"autoflow/internal/services/repo"
	"autoflow/pkg/entities/common"
	"autoflow/pkg/entities/endpoint"
	"autoflow/pkg/entities/engine"
	"context"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type ErrorService struct {
	logger *zap.SugaredLogger
	repo   *repo.Service
}

func NewErrorService(
	logger *zap.SugaredLogger,
	repo *repo.Service,
) (*ErrorService, error) {
	obj := &ErrorService{
		logger, repo,
	}

	obj.logger = obj.logger.With(zap.String("service", "endpoint_errors"))

	return obj, nil
}

func (e *ErrorService) Error(state *engine.State, res *resty.Response, resErr error) {
	e.logger.Info(
		"endpoint error",
		zap.Uint("projectId", state.GetProjectId()),
	)

	//TODO Fill values properly
	save := &endpoint.DBError{
		IDProject: common.IDProject{
			ProjectId: state.GetProjectId(),
		},
		DataError: endpoint.DataError{
			Request:    nil,
			Response:   res.Body(),
			Error:      resErr.Error(),
			StatusCode: res.StatusCode(),
			CreatedAt:  time.Now(),
		},
	}

	entity := &endpoint.DBError{}

	err := e.repo.SaveProjectObject(context.Background(), save, entity)
	if err != nil {
		e.logger.Error("failed to log", zap.Error(err))
	}
}
