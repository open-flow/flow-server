package dtos

type RunnerExecuteRequest struct {
	ActiveEvent ActiveEvent
	Raw         map[interface{}]interface{}
}

type RunnerExecuteResponse struct {
	Response map[interface{}]interface{}
}
