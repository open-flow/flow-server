package registry

import (
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/module"
)

type EndpointSource interface {
	Load() []Endpoint
}

type Endpoint interface {
	GetFunctions() []*module.FunctionDef
	GetModule() string
	Call(state *execution.State) (*execution.CallReturn, error)
	SyncUri(uri string)
	IsHealthy() bool
	Stop()
	Start()
}

type EndpointHandler interface {
}
