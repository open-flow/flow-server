package execution

import (
	"autoflow/pkg/dtos/execution"
	"autoflow/pkg/dtos/storage"
	"autoflow/pkg/orm"
	storageSvc "autoflow/pkg/services/storage"
	"context"
	"go.uber.org/zap"
	"time"
)

type ExecuteService struct {
	search    *storageSvc.SearchService
	scheduler *ScheduleService
	logger    *zap.Logger
}

func NewExecuteService(search *storageSvc.SearchService, scheduler *ScheduleService, logger *zap.Logger) *ExecuteService {
	return &ExecuteService{
		search:    search,
		scheduler: scheduler,
		logger:    logger.With(zap.String("service", "ExecuteService")),
	}
}

func (s *ExecuteService) ExecuteActiveCard(ctx context.Context, req *execution.RequestExecution) (*execution.ResponseExecution, error) {
	active, err := s.search.FindActiveGraph(ctx, req.Event)
	if err != nil {
		s.logger.Error("error finding graph", zap.Error(err))
		return nil, err
	}

	var currentVote uint64 = 0
	var responseActive *storage.ActiveGraph
	var responseActiveCard *orm.EventCard

	for _, ag := range active.Graphs {
		for _, ac := range ag.ActiveCards {
			if ac.HttpVote > currentVote {
				responseActive = ag
				responseActiveCard = ac
			}
		}
	}

	for _, ag := range active.Graphs {
		for _, ac := range ag.ActiveCards {
			if ag == responseActive && ac == responseActiveCard {
				continue
			}
			s.scheduler.Schedule(req, ag, ac, nil)
		}
	}

	if responseActive != nil {
		ch := make(chan *execution.ResponseExecution)
		s.scheduler.Schedule(req, responseActive, responseActiveCard, ch)
		select {
		case <-ctx.Done():
			return &execution.ResponseExecution{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case <-time.After(30 * time.Second):
			return &execution.ResponseExecution{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case res := <-ch:
			return res, nil
		}
	}

	return &execution.ResponseExecution{
		Scheduled: true,
	}, nil
}
