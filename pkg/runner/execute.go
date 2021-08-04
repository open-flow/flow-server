package runner

import (
	"autoflow/pkg/dtos"
	"autoflow/pkg/storage"
	"context"
	"time"
)

type ExecuteService struct {
	search    *storage.SearchService
	scheduler *ScheduleService
}

func NewExecuteService(search *storage.SearchService, scheduler *ScheduleService) *ExecuteService {
	return &ExecuteService{
		search:    search,
		scheduler: scheduler,
	}
}

func (s *ExecuteService) ExecuteActiveCard(ctx context.Context, req *dtos.ExecutionRequest) (*dtos.ExecutionResponse, error) {
	active, err := s.search.FindActiveGraph(ctx, req.Event)
	if err != nil {
		return nil, err
	}

	var withResponse *dtos.ActiveGraph
	var multiHttpResponse bool

TOP:
	for _, ag := range active.Graphs {
		for _, ac := range ag.ActiveCards {
			if ac.HttpResponse {
				if withResponse != nil {
					multiHttpResponse = true
					withResponse = nil
					break TOP
				}

				withResponse = ag
				break
			}
		}
	}

	switch multiHttpResponse {
	case true:
		for _, ag := range active.Graphs {
			s.scheduler.Schedule(req, ag)
		}
		return &dtos.ExecutionResponse{
			MultiResponseError: true,
		}, nil
	case false:
		for _, ag := range active.Graphs {
			if ag == withResponse {
				continue
			}
			s.scheduler.Schedule(req, ag)
		}

		if withResponse != nil {
			ch := s.scheduler.Schedule(req, active.Graphs[0])
			select {
			case res := <-ch:
				return res, nil
			case <-time.After(60 * time.Second):
				return &dtos.ExecutionResponse{
					Timeout: true,
					Error:   "timeout reached",
				}, nil
			}
		}

		return &dtos.ExecutionResponse{
			Scheduled: true,
		}, nil
	}

	return &dtos.ExecutionResponse{
		NoExecutions: true,
	}, nil
}
