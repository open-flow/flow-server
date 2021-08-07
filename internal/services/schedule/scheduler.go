package schedule

import (
	"autoflow/internal/services/registry"
	executionDto "autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/search"
	"go.uber.org/zap"
)

type Service struct {
	Registry *registry.Service
	logger   *zap.Logger
}

func New(regSvc *registry.Service, logger *zap.Logger) *Service {
	return &Service{
		Registry: regSvc,
		logger:   logger.With(zap.String("service", "ScheduleService")),
	}
}

func (s *Service) Schedule(
	req *executionDto.Request,
	ag *search.ActiveGraph,
	ac *graph.DBEventCard,
	res chan *executionDto.Response,
) {
	g := ag.Graph
	cursor := &executionDto.Cursor{}

	if ac.SlidePort != "" {
		cursor.Next = g.FindConnectedNodes(ac.TargetId, ac.SlidePort)
	} else {
		cursor.Next = []graph.DataConnection{
			{
				TargetId: ac.TargetId,
			},
		}
	}

	memory := &executionDto.Memory{
		Context:  req.Context,
		Response: nil,
	}

	state := &executionDto.State{
		Graph:      g,
		Card:       ac,
		RawRequest: req.RawRequest,
		Cursor:     cursor,
		Memory:     memory,
	}

	go s.run(state, res)
}
