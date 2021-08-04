package infra

import (
	"autoflow/pkg/runner"
	"autoflow/pkg/storage"
	"autoflow/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type Services struct {
	fx.In
	Search  *storage.SearchService
	Batch   *storage.BatchService
	Storage *storage.StorageService
	Execute *runner.ExecuteService
}

func NewGin(s Services) *gin.Engine {
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

	return g
}
