package execution

import (
	"autoflow/pkg/entities/graph"
)

type Request struct {
	Event      graph.DataEvent        `json:"event,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
	RawRequest interface{}            `json:"raw,omitempty"`
}

type Response struct {
	Response     interface{} `json:"response,omitempty"`
	Error        string      `json:"error,omitempty"`
	Timeout      bool        `json:"timeout,omitempty"`
	Scheduled    bool        `json:"scheduled,omitempty"`
	NoExecutions bool        `json:"noExecutions,omitempty"`
}
