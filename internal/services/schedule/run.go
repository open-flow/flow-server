package schedule

import (
	"autoflow/pkg/entities/engine"
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/utils"
	"go.uber.org/zap"
)

func (s *Service) run(state *engine.State, ch chan *engine.Response) {
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
		node := state.Graph.FindNode(current.TargetId).DataNode

		state.Cursor.Node = &node
		state.Cursor.Current = current
		state.Cursor.Next = nil

		state.Cursor.Path = append(state.Cursor.Path, current)

		wrapper, err := s.Registry.Call(state)

		logger := s.logger.With(
			zap.String("module", state.Cursor.Node.Module),
			zap.String("function", state.Cursor.Node.Function),
		)

		if err != nil {
			logger.Error("error calling function", zap.Error(err))
			ch <- &engine.Response{
				Error: "Unknown Error",
			}
			return
		}

		if wrapper.Error != nil {
			logger.Error("error response from call", zap.Any("response", wrapper))
			ch <- &engine.Response{
				Error: wrapper.Error.Message,
			}
			return
		}

		logger.Info("successful call", zap.Any("response", wrapper))

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

		var next []graph.DataConnection
		for _, sp := range slidePorts {
			next = append(next, state.Graph.FindConnectedNodes(state.Cursor.Node.LocalId, sp)...)
		}
		state.Cursor.Next = next
		state.Cursor.Node = nil
	}

	if ch != nil {
		ch <- &engine.Response{
			Response: state.Memory.Response,
		}
	}

	if len(state.Cursor.Next) > 0 {
		s.fork(state)
	}
}

func (s *Service) fork(state *engine.State) {
	for _, c := range state.Cursor.Next {
		var stateCopy engine.State
		err := utils.DeepCopy(&stateCopy, state)
		stateCopy.Cursor.Next = []graph.DataConnection{c}
		if err != nil {
			s.logger.Error("error copying state", zap.Error(err), zap.Any("state", state))
			continue
		}
		go s.run(&stateCopy, nil)
	}
}
