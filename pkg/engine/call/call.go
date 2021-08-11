package call

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Action struct {
	SlidePorts   []string               `json:"slidePorts,omitempty"`
	MergeContext map[string]interface{} `json:"mergeContext,omitempty"`
	Terminate    bool                   `json:"terminate,omitempty"`
	Freeze       string                 `json:"freeze,omitempty"`
}

type Return struct {
	Error  *Error  `json:"error,omitempty"`
	Action *Action `json:"action,omitempty"`
}