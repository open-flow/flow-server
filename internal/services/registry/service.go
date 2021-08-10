package registry

import (
	"autoflow/internal/services/endpoint"
	"autoflow/pkg/entities/engine"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Service struct {
	cache         *endpoint.Cache
	logger        *zap.SugaredLogger
	resty         *resty.Client `name:"RegistryResty"`
	endpointError *endpoint.ErrorService
}

func NewService(
	cache *endpoint.Cache,
	logger *zap.SugaredLogger,
	resty *resty.Client,
	endpointError *endpoint.ErrorService,
) (*Service, error) {
	obj := &Service{
		cache, logger, resty, endpointError,
	}

	obj.logger = obj.logger.With(zap.String("service", "registry"))

	return obj, nil
}

func (s *Service) Call(state *engine.State) (*engine.CallReturn, error) {
	container, err := s.cache.Get(state)
	if err != nil {
		return nil, err
	}
	endp, found := container.Map[state.Module()]
	if !found {
		s.logger.Error("not found", zap.String("module", state.Module()))
		return nil, engine.ErrorModuleNotFound
	}

	result := &engine.CallReturn{}
	response, err := s.resty.R().
		SetHeaders(endp.Headers).
		SetQueryParamsFromValues(endp.Values).
		SetPathParam("function", state.Function()).
		SetQueryParam("function", state.Function()).
		SetResult(result).
		Post(endp.Uri)
	if err != nil {
		s.endpointError.Error(state, response, err)
		return nil, err
	}

	return result, nil
}
