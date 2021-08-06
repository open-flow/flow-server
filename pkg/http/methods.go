package http

import (
	"autoflow/pkg/entities/batch"
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/search"
	"autoflow/pkg/entities/storage"
	"github.com/gin-gonic/gin"
)

// Call godoc
// @Summary Execute callback
// @Id Call
// @Accept  json
// @Produce  json
// @Param request body execution.Request true "request"
// @Success 200 {object} execution.Response
// @Router /call [post]
func (c *Controller) Call(g *gin.Context) {
	c.DoCall(g, c.callback.Call, func() (interface{}, error) {
		obj := execution.Request{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// ListGraph godoc
// @Summary List graphs
// @Id ListGraph
// @Accept  json
// @Produce  json
// @Param projectId query []int true "project ids"
// @Success 200 {object} storage.ListGraphResponse
// @Router /list-graphs [get]
func (c *Controller) ListGraph(g *gin.Context) {
	c.DoCall(g, c.storage.ListGraph, func() (interface{}, error) {
		obj := &storage.ListGraphRequest{}
		err := g.BindQuery(obj)
		return obj, err
	})
}

// GetGraph godoc
// @Summary Get graph with nodes, connections and event cards
// @Id GetGraph
// @Accept  json
// @Produce  json
// @Param request body storage.GetGraphRequest true "request"
// @Success 200 {object} storage.GetGraphResponse
// @Router /graph [get]
func (c *Controller) GetGraph(g *gin.Context) {
	c.DoCall(g, c.storage.GetGraph, func() (interface{}, error) {
		obj := &storage.GetGraphRequest{}
		err := g.BindQuery(obj)
		return obj, err
	})
}

// BatchSave godoc
// @Summary Save all nodes, connections and event cards
// @Id BatchSave
// @Accept  json
// @Produce  json
// @Param request body batch.SaveRequest true "request"
// @Success 200 {object} batch.SaveResponse
// @Router /batch [post]
func (c *Controller) BatchSave(g *gin.Context) {
	c.DoCall(g, c.batch.Save, func() (interface{}, error) {
		obj := &batch.SaveRequest{}
		return obj, nil
	})
}

// BatchDelete godoc
// @Summary Save all nodes, connections and event cards
// @Id BatchDelete
// @Accept  json
// @Produce  json
// @Param request body batch.DeleteRequest true "request"
// @Success 200 {object} batch.DeleteResponse
// @Router /batch [delete]
func (c *Controller) BatchDelete(g *gin.Context) {
	c.DoCall(g, c.batch.Delete, func() (interface{}, error) {
		obj := &batch.DeleteRequest{}
		return obj, nil
	})
}

// FindActive godoc
// @Summary Save all nodes, connections and event cards
// @Id FindActive
// @Accept  json
// @Produce  json
// @Param request body search.FindActiveRequest true "request"
// @Success 200 {object} search.FindActiveResponse
// @Router /find-active [post]
func (c *Controller) FindActive(g *gin.Context) {
	c.DoCall(g, c.search.FindActive, func() (interface{}, error) {
		obj := &search.FindActiveRequest{}
		return obj, nil
	})
}

// SaveGraph godoc
// @Summary Save all nodes, connections and event cards
// @Id SaveGraph
// @Accept  json
// @Produce  json
// @Param request body graph.DBGraph true "request"
// @Success 200 {object} graph.DBGraph
// @Router /graph [post]
func (c *Controller) SaveGraph(g *gin.Context) {
	c.DoCall(g, c.storage.SaveGraph, func() (interface{}, error) {
		obj := &graph.DBGraph{}
		return obj, nil
	})
}

// SaveNode godoc
// @Summary Save all nodes, connections and event cards
// @Id SaveNode
// @Accept  json
// @Produce  json
// @Param request body graph.DBNode true "request"
// @Success 200 {object} graph.DBNode
// @Router /node [post]
func (c *Controller) SaveNode(g *gin.Context) {
	c.DoCall(g, c.storage.SaveNode, func() (interface{}, error) {
		obj := &graph.DBNode{}
		return obj, nil
	})
}

// SaveConnection godoc
// @Summary Save all nodes, connections and event cards
// @Id SaveConnection
// @Accept  json
// @Produce  json
// @Param request body graph.DBConnection true "request"
// @Success 200 {object} graph.DBConnection
// @Router /connection [post]
func (c *Controller) SaveConnection(g *gin.Context) {
	c.DoCall(g, c.storage.SaveConnection, func() (interface{}, error) {
		obj := &graph.DBConnection{}
		return obj, nil
	})
}

// SaveEventCard godoc
// @Summary Save all nodes, connections and event cards
// @Id SaveEventCard
// @Accept  json
// @Produce  json
// @Param request body graph.DBEventCard true "request"
// @Success 200 {object} graph.DBEventCard
// @Router /event-card [post]
func (c *Controller) SaveEventCard(g *gin.Context) {
	c.DoCall(g, c.storage.SaveEventCard, func() (interface{}, error) {
		obj := &graph.DBEventCard{}
		return obj, nil
	})
}

// DeleteGraph godoc
// @Summary Save all nodes, connections and event cards
// @Id DeleteGraph
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /graph [delete]
func (c *Controller) DeleteGraph(g *gin.Context) {
	c.DoCall(g, c.storage.SaveGraph, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		return obj, nil
	})
}

// DeleteNode godoc
// @Summary Save all nodes, connections and event cards
// @Id DeleteNode
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /node [delete]
func (c *Controller) DeleteNode(g *gin.Context) {
	c.DoCall(g, c.storage.SaveNode, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		return obj, nil
	})
}

// DeleteConnection godoc
// @Summary Save all nodes, connections and event cards
// @Id DeleteConnection
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /connection [delete]
func (c *Controller) DeleteConnection(g *gin.Context) {
	c.DoCall(g, c.storage.SaveConnection, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		return obj, nil
	})
}

// DeleteEventCard godoc
// @Summary Save all nodes, connections and event cards
// @Id DeleteEventCard
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /event-card [delete]
func (c *Controller) DeleteEventCard(g *gin.Context) {
	c.DoCall(g, c.storage.SaveEventCard, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		return obj, nil
	})
}
