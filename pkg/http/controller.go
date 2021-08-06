package http

import (
	_ "autoflow/docs"
	"autoflow/pkg/infra"
	"autoflow/pkg/services/batch"
	"autoflow/pkg/services/callback"
	"autoflow/pkg/services/search"
	"autoflow/pkg/services/storage"
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
	engine   *gin.Engine
}

// @title Flow server
// @version 1.0

// @host localhost:8080
// @BasePath

func NewController(
	config *infra.FlowConfig,
	batchSvc *batch.Service,
	storageSvc *storage.Service,
	callbackSvc *callback.Service,
	searchSvc *search.Service,
	logger *zap.Logger,
	lc fx.Lifecycle,
) *Controller {
	e := gin.New()

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	c := &Controller{
		engine:   e,
		batch:    batchSvc,
		storage:  storageSvc,
		callback: callbackSvc,
		search:   searchSvc,
		logger:   logger.With(zap.String("service", "controller")),
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
	logger := c.logger.With(zap.String("url", g.Request.RequestURI))
	methodValue := reflect.ValueOf(method)
	obj, err := bind()
	if err != nil {
		logger.Error("binding error", zap.Error(err))
		return
	}
	resultValues := methodValue.Call([]reflect.Value{reflect.ValueOf(g), reflect.ValueOf(obj)})
	result := resultValues[0].Interface()
	errInf := resultValues[1].Interface()
	if errInf != nil {
		logger.Error("service error", zap.Error(errInf.(error)))
		g.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	g.JSON(http.StatusOK, result)
	logger.Info("request served", zap.Any("response", result))
}
