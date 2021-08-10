package schedule

import (
	"autoflow/internal/services/registry"
	"autoflow/pkg/entities/engine"
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
	req *engine.Request,
	ag *search.ActiveGraph,
	ac *graph.DBEventCard,
	res chan *engine.Response,
) {
	g := ag.Graph
	cursor := &engine.Cursor{}

	if ac.SlidePort != "" {
		cursor.Next = g.FindConnectedNodes(ac.TargetId, ac.SlidePort)
	} else {
		cursor.Next = []graph.DataConnection{
			{
				TargetId: ac.TargetId,
			},
		}
	}

	memory := &engine.Memory{
		Context:  req.Context,
		Response: nil,
	}

	state := &engine.State{
		Graph:      g,
		Card:       ac,
		RawRequest: req.RawRequest,
		Cursor:     cursor,
		Memory:     memory,
	}

	go s.run(state, res)
}
