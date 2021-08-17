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
	e.GET("/health", c.Health)

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

	loggerSvc.Info("endpoint binding complete")

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			loggerSvc.Info("starting http controller")
			errCh := make(chan error)

			go func() {
				loggerSvc.Info("listening on " + config.HttpAddr)
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
