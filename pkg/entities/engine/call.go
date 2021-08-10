package engine

type CallError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CallAction struct {
	SlidePorts   []string               `json:"slidePorts,omitempty"`
	MergeContext map[string]interface{} `json:"mergeContext,omitempty"`
	Terminate    bool                   `json:"terminate,omitempty"`
	Freeze       string                 `json:"freeze,omitempty"`
}

type CallReturn struct {
	Error  *CallError  `json:"error,omitempty"`
	Action *CallAction `json:"action,omitempty"`
}
