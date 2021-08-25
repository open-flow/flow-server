package sendpoint

import (
	"autoflow/internal/modules/serrors"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/engine/state"
	"errors"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type Registry struct {
	cache         *EndpointCache
	logger        *zap.SugaredLogger
	resty         *resty.Client `name:"RegistryResty"`
	endpointError *serrors.Errors
}

func NewRegistry(
	cache *EndpointCache,
	logger *zap.SugaredLogger,
	endpointError *serrors.Errors,
) (*Registry, error) {
	obj := &Registry{
		cache, logger, nil, endpointError,
	}

	obj.resty = resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	obj.logger = obj.logger.With(zap.String("service", "registry"))

	return obj, nil
}

func (s *Registry) Call(st *state.State) (*call.Return, error) {
	container, err := s.cache.Get(st)
	if err != nil {
		return nil, err
	}
	endp, found := container.Map[st.Cursor.Module()]
	if !found {
		s.logger.Error("not found", zap.String("module", st.Cursor.Module()))
		return nil, errors.New("module not found")
	}

	result := &call.Return{}
	response, err := s.resty.R().
		SetHeaders(endp.Headers).
		SetQueryParamsFromValues(endp.Values).
		SetPathParam("function", st.Cursor.Function()).
		SetQueryParam("function", st.Cursor.Function()).
		SetBody(st).
		SetResult(result).
		Post(endp.Uri)
	if err != nil {
		s.endpointError.Error(st, response, err)
		return nil, err
	}

	return result, nil
}
