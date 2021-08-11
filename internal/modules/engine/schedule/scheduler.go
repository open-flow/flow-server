package schedule

import (
	"autoflow/internal/modules/engine/registry"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/engine/state"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/search"
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
	req *call.Request,
	ag *search.ActiveGraph,
	ac *graph.DBEventCard,
	res chan *call.Response,
) {
	g := ag.Graph
	cursor := &state.Cursor{}

	if ac.SlidePort != "" {
		cursor.Next = g.FindConnectedNodes(ac.TargetId, ac.SlidePort)
	} else {
		cursor.Next = []graph.DataConnection{
			{
				TargetId: ac.TargetId,
			},
		}
	}

	memory := &state.Memory{
		Context:  req.Context,
		Response: nil,
	}

	st := &state.State{
		Graph:      g,
		Card:       ac,
		RawRequest: req.RawRequest,
		Cursor:     cursor,
		Memory:     memory,
	}

	go s.run(st, res)
}
