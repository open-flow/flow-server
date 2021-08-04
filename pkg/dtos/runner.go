package dtos

type ExecutionRequest struct {
	Event   *ActiveEvent           `json:"event"`
	Raw     interface{}            `json:"raw"`
	Context map[string]interface{} `json:"context"`
}

type ExecutionResponse struct {
	Response           map[interface{}]interface{} `json:"response,omitempty"`
	Error              string                      `json:"error,omitempty"`
	Timeout            bool                        `json:"timeout,omitempty"`
	Scheduled          bool                        `json:"scheduled,omitempty"`
	MultiResponseError bool                        `json:"multiHttpError,omitempty"`
	NoExecutions       bool                        `json:"noExecutions,omitempty"`
}
