package call

import (
	"autoflow/pkg/storage/graph"
)

type CallbackRequest struct {
	Event      graph.DataEvent        `json:"event,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
	RawRequest interface{}            `json:"raw,omitempty"`
}

type CallbackResponse struct {
	Response     interface{} `json:"response,omitempty"`
	Error        string      `json:"error,omitempty"`
	Timeout      bool        `json:"timeout,omitempty"`
	Scheduled    bool        `json:"scheduled,omitempty"`
	NoExecutions bool        `json:"noExecutions,omitempty"`
}
