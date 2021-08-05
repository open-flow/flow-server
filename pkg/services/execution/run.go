package execution

import (
	"autoflow/pkg/dtos/execution"
	"autoflow/pkg/utils"
	"go.uber.org/zap"
)

func (s *ScheduleService) run(state *execution.State, ch chan *execution.ResponseExecution) {
LOOP:
	for {
		nextLen := len(state.Cursor.Next)
		switch {
		case nextLen == 0:
			break LOOP
		case nextLen > 1:
			break LOOP
		}

		current := state.Cursor.Next[0]
		state.Cursor.Next = nil

		state.Cursor.Node = state.Graph.FindNode(current.TargetId)
		state.Cursor.Current = current

		state.Cursor.Path = append(state.Cursor.Path, current)

		wrapper, err := s.Registry.Call(state)

		logger := s.logger.With(
			zap.String("module", state.Cursor.Node.Module),
			zap.String("function", state.Cursor.Node.Function),
		)

		if err != nil {
			logger.Error("error calling function", zap.Error(err))
			ch <- &execution.ResponseExecution{
				Error: "Unknown Error",
			}
			return
		}

		if wrapper.Error != nil {
			logger.Error("error response from function", zap.Any("response", wrapper))
			ch <- &execution.ResponseExecution{
				Error: wrapper.Error.Message,
			}
			return
		}

		slidePorts := []string{"default"}

		if wrapper.Action != nil {
			if wrapper.Action.MergeContext != nil {
				if state.Memory.Context == nil {
					state.Memory.Context = make(map[string]interface{})
				}
				for k, v := range wrapper.Action.MergeContext {
					state.Memory.Context[k] = v
				}
			}

			if len(wrapper.Action.SlidePorts) > 0 {
				slidePorts = wrapper.Action.SlidePorts
			}
		}

		var next []*execution.Connection
		for _, sp := range slidePorts {
			next = append(next, state.Graph.FindConnectedNodes(state.Cursor.Node.LocalId, sp)...)
		}
		state.Cursor.Next = next
		state.Cursor.Node = nil
	}

	if ch != nil {
		ch <- &execution.ResponseExecution{
			Response: state.Memory.Response,
		}
	}

	if len(state.Cursor.Next) > 0 {
		s.fork(state)
	}
}

func (s *ScheduleService) fork(state *execution.State) {
	for _, c := range state.Cursor.Next {
		var stateCopy execution.State
		err := utils.DeepCopy(&stateCopy, state)
		stateCopy.Cursor.Next = []*execution.Connection{c.Copy()}
		if err != nil {
			s.logger.Error("error copying state", zap.Error(err), zap.Any("state", state))
			continue
		}
		go s.run(&stateCopy, nil)
	}
}
