package registry

import (
	"autoflow/pkg/entities/errors"
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/module"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"time"
)

type HttpEndpoint struct {
	baseUrl string
	module  string
	def     *module.ModuleDef
	c       *resty.Client
	logger  *zap.Logger
	looping bool
	exit    chan struct{}
	sync    chan struct{}
}

var _ Endpoint = (*HttpEndpoint)(nil)

func NewHttpEndpoint(baseUrl string, module string, logger *zap.Logger) *HttpEndpoint {
	client := resty.New()
	client.SetHostURL(baseUrl)
	logger = logger.With(
		zap.String("endpoint", module),
	)

	he := &HttpEndpoint{
		baseUrl: baseUrl,
		c:       client,
		logger:  logger,
		module:  module,
		sync:    make(chan struct{}),
		exit:    make(chan struct{}),
	}

	return he
}

func (h *HttpEndpoint) panicUnhealthy() {
	if h.def == nil {
		panic("Working with uninitialized HttpEndpoint")
	}
}

func (h *HttpEndpoint) Initialize() error {
	var moduleDef module.ModuleDef
	h.c.SetHostURL(h.baseUrl)

	_, err := h.c.R().SetResult(&moduleDef).Get("")

	if err != nil {
		h.logger.Error("unable to initialize", zap.Error(err))
		h.def = nil
		return err
	}

	h.logger.Info("successful initialization",
		zap.String("url", h.baseUrl),
	)

	h.def = &moduleDef
	return nil
}

func (h *HttpEndpoint) GetFunctions() []*module.FunctionDef {
	h.panicUnhealthy()
	return h.def.Functions
}

func (h *HttpEndpoint) GetModule() string {
	return h.module
}

func (h *HttpEndpoint) Call(state *execution.State) (*execution.CallReturn, error) {
	h.panicUnhealthy()
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

func (h *HttpEndpoint) SyncUri(uri string) {
	if h.baseUrl == uri && h.def != nil {
		return
	}

	h.baseUrl = uri
}

func (h *HttpEndpoint) IsHealthy() bool {
	return h.def != nil
}

func (h *HttpEndpoint) Stop() {
	h.looping = false
	h.exit <- struct{}{}
	h.logger.Info("stopping")
}

func (h *HttpEndpoint) Start() {
	if h.looping {
		return
	}
	h.logger.Info("starting")
	h.looping = true
	go h.syncLoop()
	h.sync <- struct{}{}
}

func (h *HttpEndpoint) syncLoop() {
	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-h.exit:
			return
		case <-tick.C:
			if !h.IsHealthy() {
				_ = h.Initialize()
			}
		case <-h.sync:
			_ = h.Initialize()
		}
	}
}
