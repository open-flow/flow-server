package registry

import (
	endpoint2 "autoflow/internal/modules/storage/endpoint"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/engine/state"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type Service struct {
	cache         *endpoint2.Cache
	logger        *zap.SugaredLogger
	resty         *resty.Client `name:"RegistryResty"`
	endpointError *endpoint2.ErrorService
}

func NewService(
	cache *endpoint2.Cache,
	logger *zap.SugaredLogger,
	resty *resty.Client,
	endpointError *endpoint2.ErrorService,
) (*Service, error) {
	obj := &Service{
		cache, logger, resty, endpointError,
	}

	obj.logger = obj.logger.With(zap.String("service", "registry"))

	return obj, nil
}

func (s *Service) Call(st *state.State) (*call.Return, error) {
	container, err := s.cache.Get(st)
	if err != nil {
		return nil, err
	}
	endp, found := container.Map[st.Module()]
	if !found {
		s.logger.Error("not found", zap.String("module", st.Module()))
		return nil, state.ErrorModuleNotFound
	}

	result := &call.Return{}
	response, err := s.resty.R().
		SetHeaders(endp.Headers).
		SetQueryParamsFromValues(endp.Values).
		SetPathParam("function", st.Function()).
		SetQueryParam("function", st.Function()).
		SetResult(result).
		Post(endp.Uri)
	if err != nil {
		s.endpointError.Error(st, response, err)
		return nil, err
	}

	return result, nil
}
