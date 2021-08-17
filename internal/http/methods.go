package http

import (
	"autoflow/pkg/common"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/storage/batch"
	"autoflow/pkg/storage/endpoint"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/search"
	"autoflow/pkg/storage/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

//@router /callback [post]
//@param request body call.CallbackRequest true "callback"
//@success 200 {object} call.CallbackResponse
func (c *Controller) Callback(g *gin.Context) {
	req := &call.CallbackRequest{}
	err := g.Bind(req)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.callback.Call(g, req)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

func (c *Controller) Health(g *gin.Context) {
	g.String(200, "ok")
}

//@router /graph/list [get]
//@param projectId query []int true "project ids"
//@success 200 {object} storage.ListGraphResponse
func (c *Controller) ListGraph(g *gin.Context) {
	obj := &storage.ListGraphRequest{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.ListGraph(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /graph [get]
//@param projectId query int true "project id"
//@param graphId query int true "graph id"
//@success 200 {object} graph.DBGraph
func (c *Controller) GetGraph(g *gin.Context) {
	obj := &common.ProjectModel{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.GetGraph(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /batch [post]
//@param body body batch.SaveRequest true "save request"
//@success 200 {object} batch.SaveResponse
func (c *Controller) BatchSave(g *gin.Context) {
	obj := &batch.SaveRequest{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.batch.Save(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /batch [delete]
//@param body body batch.DeleteRequest true "delete request"
//@success 200 {string} string "ok"
func (c *Controller) BatchDelete(g *gin.Context) {
	obj := &batch.DeleteRequest{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.batch.Delete(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}

//@router /find-active [post]
//@param body body search.FindActiveRequest true "active event"
//@success 200 {object} search.FindActiveResponse
func (c *Controller) FindActive(g *gin.Context) {
	obj := &search.FindActiveRequest{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.search.FindActive(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /graph [post]
//@param body body graph.DBGraph true "graph"
//@success 200 {object} graph.DBGraph
func (c *Controller) SaveGraph(g *gin.Context) {
	obj := &graph.DBGraph{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.SaveGraph(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /node [post]
//@param body body graph.DBNode true "node"
//@success 200 {object} graph.DBNode
func (c *Controller) SaveNode(g *gin.Context) {
	obj := &graph.DBNode{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.SaveNode(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /connection [post]
//@param body body graph.DBConnection true "connection"
//@success 200 {object} graph.DBConnection
func (c *Controller) SaveConnection(g *gin.Context) {
	obj := &graph.DBConnection{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.SaveConnection(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /event-card [post]
//@param body body graph.DBEventCard true "connection"
//@success 200 {object} graph.DBEventCard
func (c *Controller) SaveEventCard(g *gin.Context) {
	obj := &graph.DBEventCard{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.storage.SaveEventCard(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /graph [delete]
//@param body body common.ProjectModel true "graph id"
//@success 200 {string} string "ok"
func (c *Controller) DeleteGraph(g *gin.Context) {
	obj := &common.ProjectModel{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.storage.DeleteGraph(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}

//@router /node [delete]
//@param body body graph.GraphObject true "node id"
//@success 200 {string} string "ok"
func (c *Controller) DeleteNode(g *gin.Context) {
	obj := &graph.GraphObject{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.storage.DeleteNode(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}

//@router /connection [delete]
//@param body body graph.GraphObject true "connection id"
//@success 200 {string} string "ok"
func (c *Controller) DeleteConnection(g *gin.Context) {
	obj := &graph.GraphObject{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.storage.DeleteConnection(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}

//@router /event-card [delete]
//@param body body graph.GraphObject true "event card id"
//@success 200 {string} string "ok"
func (c *Controller) DeleteEventCard(g *gin.Context) {
	obj := &graph.GraphObject{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.storage.DeleteEventCard(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}

//@router /endpoint/list [get]
//@param projectId query int true "project id"
//@success 200 {object} endpoint.Container
func (c *Controller) ListEndpoint(g *gin.Context) {
	obj := &common.ProjectSpace{}
	err := g.BindQuery(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.endpoint.List(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /endpoint [post]
//@param body body endpoint.DBEndpoint true "endpoint"
//@success 200 {object} endpoint.DBEndpoint
func (c *Controller) SaveEndpoint(g *gin.Context) {
	obj := &endpoint.DBEndpoint{}
	err := g.BindJSON(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	res, err := c.endpoint.Save(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, res)
}

//@router /endpoint [delete]
//@param body body common.ProjectModel true "endpoint id"
//@success 200 {string} string "ok"
func (c *Controller) DeleteEndpoint(g *gin.Context) {
	obj := &common.ProjectModel{}
	err := g.Bind(obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	err = c.endpoint.Delete(g, obj)
	if err != nil {
		_ = g.Error(err)
		return
	}
	g.JSON(http.StatusOK, "ok")
}
