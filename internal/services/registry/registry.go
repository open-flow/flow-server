package registry

import (
	"autoflow/internal/services/module"
	"autoflow/pkg/entities/errors"
	"autoflow/pkg/entities/execution"
	modulesDto "autoflow/pkg/entities/module"
	"autoflow/pkg/topics"
	"context"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	Static         map[string]Endpoint
	Dynamic        map[string]Endpoint
	mut            *sync.RWMutex
	logger         *zap.Logger
	pristineLogger *zap.Logger
	modules        *module.Service
}

func NewRegistryService(
	logger *zap.Logger,
	lc fx.Lifecycle,
	modules *module.Service,
	nc *nats.Conn,
) *Service {
	service := &Service{
		Static:         map[string]Endpoint{},
		Dynamic:        map[string]Endpoint{},
		modules:        modules,
		mut:            &sync.RWMutex{},
		logger:         logger.With(zap.String("service", "RegistryService")),
		pristineLogger: logger,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			service.handleSync()
			_, err := nc.Subscribe(topics.MODULES_SYNC, func(msg *nats.Msg) {
				service.handleSync()
			})
			return err
		},
	})

	return service
}

func (r *Service) RegisterEndpoint(endpoint Endpoint) {
	r.Static[endpoint.GetModule()] = endpoint
}

func (r *Service) Call(state *execution.State) (*execution.CallReturn, error) {
	r.mut.RLock()
	moduleStr := state.Cursor.Node.Module
	endpointStatic, foundStatic := r.Static[moduleStr]
	endpointDynamic, foundDynamic := r.Dynamic[moduleStr]
	r.mut.RUnlock()

	var endpoint Endpoint

	switch {
	case foundStatic:
		endpoint = endpointStatic
	case foundDynamic:
		endpoint = endpointDynamic
	default:
		r.logger.Error("endpoint not found", zap.String("module", moduleStr))
		return &execution.CallReturn{
			Error: &execution.CallError{
				Code:    "Module not found",
				Message: "Module not found",
			},
		}, nil
	}

	if endpoint.IsHealthy() {
		return endpoint.Call(state)
	}

	return nil, errors.EndpointUnhealthy
}

func (r *Service) handleSync() {
	res, err := r.modules.List(context.Background(), &modulesDto.ListRequest{})
	if err != nil {
		r.logger.Error("unable to list modules", zap.Error(err))
		return
	}

	newDynamic := make(map[string]Endpoint)

	for _, mod := range res.Modules {
		end, found := r.Dynamic[mod.Module]
		if found {
			end.SyncUri(mod.Uri)
		} else {
			end = NewHttpEndpoint(mod.Uri, mod.Module, r.pristineLogger)
			end.Start()
		}
		newDynamic[end.GetModule()] = end
	}

	r.mut.Lock()
	oldDynamic := r.Dynamic
	r.Dynamic = newDynamic
	r.mut.Unlock()

	r.mut.RLock()
	defer r.mut.RUnlock()
	for k, v := range oldDynamic {
		if _, found := newDynamic[k]; !found {
			v.Stop()
		}
	}
}
