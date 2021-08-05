package infra

import (
	runner2 "autoflow/pkg/services/execution"
	storage2 "autoflow/pkg/services/storage"
	"autoflow/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Services struct {
	fx.In
	Config  *FlowConfig
	Search  *storage2.SearchService
	Batch   *storage2.BatchService
	Storage *storage2.GraphService
	Execute *runner2.ExecuteService
}

func NewGin(ls fx.Lifecycle, s Services) *gin.Engine {
	g := gin.New()
	bg := utils.BindGinFactory(g)

	bg(http.MethodPost, "search/find-active-graph", s.Search.FindActiveGraph)

	bg(http.MethodPost, "batch/save", s.Batch.Save)
	bg(http.MethodPost, "batch/delete", s.Batch.Delete)

	bg(http.MethodPost, "graph/delete-graph", s.Storage.DeleteGraph)
	bg(http.MethodPost, "graph/delete-node", s.Storage.DeleteNode)
	bg(http.MethodPost, "graph/delete-connection", s.Storage.DeleteConnection)
	bg(http.MethodPost, "graph/delete-event-card", s.Storage.DeleteEventCard)

	bg(http.MethodPost, "graph/save-graph", s.Storage.SaveGraph)
	bg(http.MethodPost, "graph/save-connection", s.Storage.SaveConnection)
	bg(http.MethodPost, "graph/save-node", s.Storage.SaveNode)
	bg(http.MethodPost, "graph/save-event-card", s.Storage.SaveEventCard)

	bg(http.MethodGet, "graph/list-graphs", s.Storage.ListGraph)
	bg(http.MethodGet, "graph/full-graph", s.Storage.GetFullGraph)
	bg(http.MethodPost, "execute", s.Execute.ExecuteActiveCard)

	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			errCh := make(chan error)

			go func() {
				err := g.Run(s.Config.HttpAddr)
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

	return g
}
