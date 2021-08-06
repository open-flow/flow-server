package registry

import (
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/registry"
)

type EndpointSource interface {
	Load() []Endpoint
}

type Endpoint interface {
	GetFunctions() []*registry.FunctionDef
	GetModule() string
	Call(state *execution.State) (*execution.CallReturn, error)
}

type EndpointHandler interface {
}
