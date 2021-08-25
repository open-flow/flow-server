package sscheduler

import (
	"autoflow/pkg/engine/call"
	"autoflow/pkg/engine/state"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/utils"
	"go.uber.org/zap"
)

func (s *Schedule) run(st *state.State, ch chan *call.CallbackResponse) {
LOOP:
	for {
		nextLen := len(st.Cursor.Next)
		switch {
		case nextLen == 0:
			break LOOP
		case nextLen > 1:
			break LOOP
		}

		loopExit := s.loopConstrain(st)
		if loopExit {
			s.logger.Error("exiting due to loop constrain",
				zap.Uint("graphId", st.Graph.Id),
				zap.Uint("projectId", st.Graph.ProjectId),
				zap.Any("path", st.Cursor.Path),
			)
			break LOOP
		}

		current := st.Cursor.Next[0]
		node := st.Graph.FindNode(current.TargetId).DataNode

		st.Cursor.Node = &node
		st.Cursor.Current = current
		st.Cursor.Next = nil

		st.Cursor.Path = append(st.Cursor.Path, current)

		wrapper, err := s.Registry.Call(st)

		logger := s.logger.With(
			zap.String("module", st.Cursor.Node.Module),
			zap.String("function", st.Cursor.Node.Function),
		)

		if err != nil {
			logger.Error("error calling function", zap.Error(err))
			ch <- &call.CallbackResponse{
				Error: "Unknown Error",
			}
			return
		}

		if wrapper.Error != nil {
			logger.Error("error response from call", zap.Any("response", wrapper))
			ch <- &call.CallbackResponse{
				Error: wrapper.Error.Message,
			}
			return
		}

		logger.Info("successful call", zap.Any("response", wrapper))

		slidePorts := []string{"default"}

		if wrapper.Action != nil {
			if wrapper.Action.MergeContext != nil {
				if st.Memory.Context == nil {
					st.Memory.Context = make(map[string]interface{})
				}
				for k, v := range wrapper.Action.MergeContext {
					st.Memory.Context[k] = v
				}
			}

			if len(wrapper.Action.SlidePorts) > 0 {
				slidePorts = wrapper.Action.SlidePorts
			}

			if wrapper.Action.Response != nil {
				st.Memory.Response = wrapper.Action.Response
			}
		}

		var next []graph.DataConnection
		for _, sp := range slidePorts {
			next = append(next, st.Graph.FindConnectedNodes(st.Cursor.Node.LocalId, sp)...)
		}
		st.Cursor.Next = next
		st.Cursor.Node = nil
	}

	if ch != nil {
		ch <- &call.CallbackResponse{
			Response: st.Memory.Response,
		}
	}

	if len(st.Cursor.Next) > 0 {
		s.fork(st)
	}
}

func (s *Schedule) loopConstrain(st *state.State) bool {
	m := make(map[uint]uint)
	for _, v := range st.Cursor.Path {
		cnt, _ := m[v.LocalID]
		cnt++
		if cnt > state.LOOPING_MAX_COUNTER {
			return true
		}
		m[v.LocalID] = cnt
	}
	return false
}

func (s *Schedule) fork(st *state.State) {
	for _, c := range st.Cursor.Next {
		var stateCopy state.State
		err := utils.DeepCopy(&stateCopy, st)
		stateCopy.Cursor.Next = []graph.DataConnection{c}
		if err != nil {
			s.logger.Error("error copying state", zap.Error(err), zap.Any("state", st))
			continue
		}
		go s.run(&stateCopy, nil)
	}
}
