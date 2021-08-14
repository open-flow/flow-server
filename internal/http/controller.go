package http

import (
	"autoflow/docs"
	"autoflow/internal/infra"
	"autoflow/internal/modules/scallback"
	"autoflow/internal/modules/sendpoint"
	"autoflow/internal/modules/sgraph"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"reflect"
	"time"
)

type Controller struct {
	batch    *sgraph.GraphBatch
	storage  *sgraph.Graph
	callback *scallback.Callback
	search   *sgraph.Active
	logger   *zap.Logger
	gin      *gin.Engine
	endpoint *sendpoint.Endpoint
}

func NewController(
	batchSvc *sgraph.GraphBatch,
	storageSvc *sgraph.Graph,
	callbackSvc *scallback.Callback,
	searchSvc *sgraph.Active,
	loggerSvc *zap.Logger,
	endpointSvc *sendpoint.Endpoint,
	config *infra.FlowConfig,
	lc fx.Lifecycle,
) *Controller {
	e := gin.New()

	docs.SwaggerInfo.Host = config.HostName
	docs.SwaggerInfo.Version = "0.0.1"

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c := &Controller{
		gin:      e,
		batch:    batchSvc,
		storage:  storageSvc,
		callback: callbackSvc,
		search:   searchSvc,
		logger:   loggerSvc.With(zap.String("service", "controller")),
		endpoint: endpointSvc,
	}

	e.POST("/callback", c.Callback)

	e.GET("/graph/list", c.ListGraph)
	e.GET("/graph", c.GetGraph)
	e.POST("/graph", c.SaveGraph)
	e.DELETE("/graph", c.DeleteGraph)

	e.POST("/node", c.SaveNode)
	e.DELETE("/node", c.DeleteNode)

	e.POST("/event-card", c.SaveEventCard)
	e.DELETE("/event-card", c.DeleteEventCard)

	e.POST("/connection", c.SaveConnection)
	e.DELETE("/connection", c.DeleteConnection)

	e.POST("/batch", c.BatchSave)
	e.DELETE("/batch", c.BatchDelete)

	e.POST("/find-active", c.FindActive)

	e.GET("/endpoint/list", c.ListEndpoint)
	e.POST("/endpoint", c.SaveEndpoint)
	e.DELETE("/endpoint", c.DeleteEndpoint)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errCh := make(chan error)

			go func() {
				err := e.Run(config.HttpAddr)
				errCh <- err
			}()

			select {
			case err := <-errCh:
				return err
			case <-time.After(1 * time.Second):
				return nil
			}
		},
	})

	return c
}

func (c *Controller) DoCall(g *gin.Context, method interface{}, bind func() (interface{}, error)) {
	logger := c.logger.With(zap.String("url", g.Request.RequestURI), zap.String("method", g.Request.Method))
	methodValue := reflect.ValueOf(method)
	obj, err := bind()
	if err != nil {
		logger.Error("binding error", zap.Error(err))
		return
	}
	resultValues := methodValue.Call([]reflect.Value{reflect.ValueOf(g), reflect.ValueOf(obj)})
	var result interface{}
	var errInf interface{}

	switch len(resultValues) {
	case 1:
		errInf = resultValues[0].Interface()
	case 2:
		result = resultValues[0].Interface()
		errInf = resultValues[1].Interface()
	default:
		panic("")
	}

	if errInf != nil {
		err := errInf.(error)
		logger.Error("service error", zap.Error(err))
		_ = g.Error(err)
		g.JSON(http.StatusInternalServerError, HttpError{
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, result)
	logger.Info("request served", zap.Any("response", result))
}
