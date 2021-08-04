package runner

import "autoflow/pkg/dtos"

type ScheduleService struct {
}

func NewScheduleService() *ScheduleService {
	return &ScheduleService{}
}

func (s *ScheduleService) Schedule(req *dtos.ExecutionRequest, active *dtos.ActiveGraph) chan *dtos.ExecutionResponse {
	ch := make(chan *dtos.ExecutionResponse)

	if len(active.ActiveCards) > 1 {
		//TODO warning
	}

	card := active.ActiveCards[0]
	cursor := &dtos.ExecutionCursor{
		EventCard:  card,
		Graph:      active.Graph,
		Visited:    make([]dtos.Visited, 0),
		TargetId:   card.TargetId,
		Context:    req.Context,
		RawRequest: req.Raw,
	}

	go s.run(cursor, ch)
	return ch
}
