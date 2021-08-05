package registry

import (
	"autoflow/pkg/dtos/execution"
	"go.uber.org/zap"
)

type RegistryService struct {
	Endpoints map[string]Endpoint
	logger    *zap.Logger
}

func NewRegistryService(logger *zap.Logger) *RegistryService {
	service := &RegistryService{
		Endpoints: map[string]Endpoint{},
		logger:    logger.With(zap.String("service", "RegistryService")),
	}
	return service
}

func (r *RegistryService) RegisterEndpoint(endpoint Endpoint) {
	r.Endpoints[endpoint.GetModule()] = endpoint
}

func (r *RegistryService) Call(state *execution.State) (*execution.CallReturn, error) {
	module := state.Cursor.Node.Module
	endpoint, found := r.Endpoints[module]
	if !found {
		r.logger.Error("endpoint not found", zap.String("module", module))
		return &execution.CallReturn{
			Error: &execution.CallError{
				Code:    "Module not found",
				Message: "Module not found",
			},
		}, nil
	}

	return endpoint.Call(state)
}
