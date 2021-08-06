package registry

import (
	"autoflow/pkg/entities/errors"
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/registry"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

type HttpEndpoint struct {
	BaseUrl string
	Module  *registry.ModuleDef
	c       *resty.Client
	logger  *zap.Logger
}

var _ Endpoint = (*HttpEndpoint)(nil)

func NewHttpEndpoint(baseUrl string, module string, logger *zap.Logger) *HttpEndpoint {
	client := resty.New()
	client.SetHostURL(baseUrl)
	logger = logger.With(
		zap.String("HttpEndpoint", module),
	)

	return &HttpEndpoint{
		BaseUrl: baseUrl,
		c:       client,
		logger:  logger,
	}
}

func (h *HttpEndpoint) panicUninitialized() {
	if h.Module == nil {
		panic("Working with uninitialized HttpEndpoint")
	}
}

func (h *HttpEndpoint) Initialize() error {
	var moduleDef registry.ModuleDef
	_, err := h.c.R().SetResult(&moduleDef).Get("")

	if err != nil {
		h.logger.Error("initialization error", zap.Error(err))
		h.Module = nil
		return err
	}

	h.Module = &moduleDef
	return nil
}

func (h *HttpEndpoint) GetFunctions() []*registry.FunctionDef {
	h.panicUninitialized()
	return h.Module.Functions
}

func (h *HttpEndpoint) GetModule() string {
	h.panicUninitialized()
	return h.Module.Module
}

func (h *HttpEndpoint) Call(state *execution.State) (*execution.CallReturn, error) {
	var result execution.CallReturn
	function := state.Cursor.Node.Function

	res, err := h.c.R().
		SetBody(state).
		SetResult(&result).
		Post(function)

	if err != nil {
		h.logger.Error("calling error", zap.Error(err))
		return nil, err
	}

	if res.StatusCode() >= 300 || res.StatusCode() < 200 {
		err = errors.UnknownHttpResponse
		h.logger.Error(
			"unknown response",
			zap.String("function", function),
			zap.ByteString("response", res.Body()),
		)
		return nil, err
	}

	return &result, nil
}
