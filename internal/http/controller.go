package http

import (
	_ "autoflow/docs"
	"autoflow/internal/infra"
	"autoflow/internal/modules/engine/callback"
	"autoflow/internal/modules/storage/batch"
	"autoflow/internal/modules/storage/endpoint"
	"autoflow/internal/modules/storage/search"
	"autoflow/internal/modules/storage/storage"
	"autoflow/pkg/common"
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
	batch    *batch.Service
	storage  *storage.Service
	callback *callback.Service
	search   *search.Service
	logger   *zap.Logger
	gin      *gin.Engine
	endpoint *endpoint.Controller
}

// @title Flow server
// @version 1.0

// @host localhost:8080
// @BasePath

func NewController(
	batchSvc *batch.Service,
	storageSvc *storage.Service,
	callbackSvc *callback.Service,
	searchSvc *search.Service,
	loggerSvc *zap.Logger,
	endpointSvc *endpoint.Controller,
	config *infra.FlowConfig,
	lc fx.Lifecycle,
) *Controller {
	e := gin.New()

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

	e.POST("/call", c.Call)

	e.GET("/list-graphs", c.ListGraph)
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

	e.GET("/list-module", c.ListEndpoint)
	e.POST("/module", c.SaveEndpoint)
	e.DELETE("/module", c.DeleteEndpoint)

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
		g.JSON(http.StatusInternalServerError, common.HttpError{
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, result)
	logger.Info("request served", zap.Any("response", result))
}
