package registry

import (
	"autoflow/pkg/dtos/execution"
	"autoflow/pkg/dtos/registry"
)

type EndpointSource interface {
	Load() []Endpoint
}

type Endpoint interface {
	GetFunctions() []*registry.EndpointFunctionDef
	GetModule() string
	Call(state *execution.State) (*execution.CallReturn, error)
}

type EndpointHandler interface {
}
