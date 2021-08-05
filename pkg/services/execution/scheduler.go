package execution

import (
	"autoflow/pkg/dtos/execution"
	"autoflow/pkg/dtos/storage"
	"autoflow/pkg/orm"
	"autoflow/pkg/services/registry"
	"go.uber.org/zap"
)

type ScheduleService struct {
	Registry *registry.RegistryService
	logger   *zap.Logger
}

func NewScheduleService(regSvc *registry.RegistryService, logger *zap.Logger) *ScheduleService {
	return &ScheduleService{
		Registry: regSvc,
		logger:   logger.With(zap.String("service", "ScheduleService")),
	}
}

func (s *ScheduleService) Schedule(
	req *execution.RequestExecution,
	ag *storage.ActiveGraph,
	ac *orm.EventCard,
	res chan *execution.ResponseExecution,
) {
	graph := ag.Graph

	activation := &execution.Activation{
		Card:       ac,
		RawRequest: req.RawRequest,
	}
	cursor := &execution.Cursor{
		Node: graph.FindNode(ac.SourceLocalId),
	}
	if ac.SlidePort != "" {
		cursor.Next = graph.FindConnectedNodes(ac.SourceLocalId, ac.SlidePort)
	}

	memory := &execution.Memory{
		Context:  req.Context,
		Response: nil,
	}

	state := &execution.State{
		Activation: activation,
		Cursor:     cursor,
		Graph:      graph,
		Memory:     memory,
	}

	go s.run(state, res)
}
