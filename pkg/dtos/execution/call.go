package execution

type CallError struct {
	Code    string
	Message string
}

type CallAction struct {
	SlidePorts   []string
	MergeContext map[string]interface{}
	Terminate    bool
	Freeze       string
}

type CallReturn struct {
	Error  *CallError
	Action *CallAction
}
