package execution

import (
	"autoflow/pkg/entities/graph"
)

type Request struct {
	Event      graph.DataEvent        `json:"event"`
	Context    map[string]interface{} `json:"context"`
	RawRequest interface{}            `json:"raw"`
}

type Response struct {
	Response     interface{} `json:"response,omitempty"`
	Error        string      `json:"error,omitempty"`
	Timeout      bool        `json:"timeout,omitempty"`
	Scheduled    bool        `json:"scheduled,omitempty"`
	NoExecutions bool        `json:"noExecutions,omitempty"`
}
