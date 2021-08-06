package registry

import (
	execution2 "autoflow/pkg/entities/execution"
	registry2 "autoflow/pkg/entities/registry"
)

type EndpointSource interface {
	Load() []Endpoint
}

type Endpoint interface {
	GetFunctions() []*registry2.FunctionDef
	GetModule() string
	Call(state *execution2.State) (*execution2.CallReturn, error)
}

type EndpointHandler interface {
}
