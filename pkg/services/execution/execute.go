package execution

import (
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/graph"
	searchDto "autoflow/pkg/entities/search"
	"autoflow/pkg/services/schedule"
	"autoflow/pkg/services/search"
	"context"
	"go.uber.org/zap"
	"time"
)

type ExecuteService struct {
	search    *search.Service
	scheduler *schedule.Service
	logger    *zap.Logger
}

func NewExecuteService(search *search.Service, scheduler *schedule.Service, logger *zap.Logger) *ExecuteService {
	return &ExecuteService{
		search:    search,
		scheduler: scheduler,
		logger:    logger.With(zap.String("service", "ExecuteService")),
	}
}

func (s *ExecuteService) ExecuteActiveCard(ctx context.Context, req *execution.Request) (*execution.Response, error) {
	active, err := s.search.FindActive(ctx, &searchDto.FindActiveRequest{
		DataEvent: req.Event,
	})

	if err != nil {
		s.logger.Error("error finding graph", zap.Error(err))
		return nil, err
	}

	var currentVote uint = 0
	var responseActive *searchDto.ActiveGraph
	var responseActiveCard *graph.DBEventCard

	for _, ag := range active.Graphs {
		for _, ac := range ag.Active {
			if ac.HttpVote > currentVote {
				responseActive = ag
				responseActiveCard = ac
			}
		}
	}

	for _, ag := range active.Graphs {
		for _, ac := range ag.Active {
			if ag == responseActive && ac == responseActiveCard {
				continue
			}
			s.scheduler.Schedule(req, ag, ac, nil)
		}
	}

	if responseActive != nil {
		ch := make(chan *execution.Response)
		s.scheduler.Schedule(req, responseActive, responseActiveCard, ch)
		select {
		case <-ctx.Done():
			return &execution.Response{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case <-time.After(30 * time.Second):
			return &execution.Response{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case res := <-ch:
			return res, nil
		}
	}

	return &execution.Response{
		Scheduled: true,
	}, nil
}
