package callback

import (
	"autoflow/internal/modules/engine/schedule"
	"autoflow/internal/modules/storage/search"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/storage/graph"
	searchDto "autoflow/pkg/storage/search"
	"context"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	search    *search.Service
	scheduler *schedule.Service
	logger    *zap.Logger
}

func New(search *search.Service, scheduler *schedule.Service, logger *zap.Logger) *Service {
	return &Service{
		search:    search,
		scheduler: scheduler,
		logger:    logger.With(zap.String("service", "callback")),
	}
}

func (s *Service) Call(ctx context.Context, req *call.Request) (*call.Response, error) {
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
		ch := make(chan *call.Response)
		s.scheduler.Schedule(req, responseActive, responseActiveCard, ch)
		select {
		case <-ctx.Done():
			return &call.Response{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case <-time.After(30 * time.Second):
			return &call.Response{
				Timeout: true,
				Error:   "timeout reached",
			}, nil
		case res := <-ch:
			return res, nil
		}
	}

	return &call.Response{
		Scheduled: true,
	}, nil
}
