package execution

import "autoflow/pkg/dtos/storage"

type RequestExecution struct {
	Event      *storage.RequestFindActiveGraph `json:"event"`
	Context    map[string]interface{}          `json:"context"`
	RawRequest interface{}                     `json:"raw"`
}

type ResponseExecution struct {
	Response     interface{} `json:"response,omitempty"`
	Error        string      `json:"error,omitempty"`
	Timeout      bool        `json:"timeout,omitempty"`
	Scheduled    bool        `json:"scheduled,omitempty"`
	NoExecutions bool        `json:"noExecutions,omitempty"`
}
