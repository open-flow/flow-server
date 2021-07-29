package binding

import (
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/search"
	"autoflow/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func BindStorageGin(e *gin.Engine, db *gorm.DB) {
	searchSvc := search.NewService(db)
	batchSvc := graph.NewBatchService(db)
	graphSvc := graph.NewGraphService(db)

	bg := utils.BindGinFactory(e)

	bg(http.MethodPost, "search/find-active-graph", searchSvc.FindActiveGraph)

	bg(http.MethodPost, "batch/save", batchSvc.Save)
	bg(http.MethodPost, "batch/delete", batchSvc.Delete)

	bg(http.MethodPost, "graph/delete-graph", graphSvc.DeleteGraph)
	bg(http.MethodPost, "graph/delete-node", graphSvc.DeleteNode)
	bg(http.MethodPost, "graph/delete-connection", graphSvc.DeleteConnection)
	bg(http.MethodPost, "graph/delete-event-card", graphSvc.DeleteEventCard)

	bg(http.MethodPost, "graph/save-graph", graphSvc.SaveGraph)
	bg(http.MethodPost, "graph/save-connection", graphSvc.SaveConnection)
	bg(http.MethodPost, "graph/save-node", graphSvc.SaveNode)
	bg(http.MethodPost, "graph/save-event-card", graphSvc.SaveEventCard)

	bg(http.MethodGet, "graph/list-graphs", graphSvc.ListGraph)
	bg(http.MethodGet, "graph/full-graph", graphSvc.GetFullGraph)
}
