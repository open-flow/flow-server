package serrors

import (
	"autoflow/internal/modules/srepo"
	"autoflow/pkg/common"
	"autoflow/pkg/engine/state"
	"autoflow/pkg/storage/endpoint"
	"context"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type Errors struct {
	logger *zap.SugaredLogger
	repo   *srepo.Repo
}

func NewErrors(
	logger *zap.SugaredLogger,
	repo *srepo.Repo,
) (*Errors, error) {
	obj := &Errors{
		logger, repo,
	}

	obj.logger = obj.logger.With(zap.String("service", "endpoint_errors"))

	return obj, nil
}

func (e *Errors) Error(state *state.State, res *resty.Response, resErr error) {
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
