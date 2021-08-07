package http

import (
	"autoflow/pkg/entities/batch"
	"autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/module"
	"autoflow/pkg/entities/search"
	"autoflow/pkg/entities/storage"
	"github.com/gin-gonic/gin"
)

// Call godoc
// @Summary executes callback
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
// @Summary lists graphs
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
// @Summary gets graph with Nodes, Connections and EventCards
// @Id GetGraph
// @Accept  json
// @Produce  json
// @Param projectId query int true "project id"
// @Param id query int true "graph id"
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
// @Summary saves new Nodes, Connections and EventCards
// @Id BatchSave
// @Accept  json
// @Produce  json
// @Param request body batch.SaveRequest true "request"
// @Success 200 {object} batch.SaveResponse
// @Router /batch [post]
func (c *Controller) BatchSave(g *gin.Context) {
	c.DoCall(g, c.batch.Save, func() (interface{}, error) {
		obj := &batch.SaveRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// BatchDelete godoc
// @Summary deletes all Nodes, Connections and EventCards from the graph
// @Id BatchDelete
// @Accept  json
// @Produce  json
// @Param request body batch.DeleteRequest true "request"
// @Success 200 {object} batch.DeleteResponse
// @Router /batch [delete]
func (c *Controller) BatchDelete(g *gin.Context) {
	c.DoCall(g, c.batch.Delete, func() (interface{}, error) {
		obj := &batch.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// FindActive godoc
// @Summary finds graphs that would activate for event
// @Id FindActive
// @Accept  json
// @Produce  json
// @Param request body search.FindActiveRequest true "request"
// @Success 200 {object} search.FindActiveResponse
// @Router /find-active [post]
func (c *Controller) FindActive(g *gin.Context) {
	c.DoCall(g, c.search.FindActive, func() (interface{}, error) {
		obj := &search.FindActiveRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// SaveGraph godoc
// @Summary saves or updates graph. Nodes, Connections and EventCards are ignored
// @Id SaveGraph
// @Accept  json
// @Produce  json
// @Param request body graph.DBGraph true "request"
// @Success 200 {object} graph.DBGraph
// @Router /graph [post]
func (c *Controller) SaveGraph(g *gin.Context) {
	c.DoCall(g, c.storage.SaveGraph, func() (interface{}, error) {
		obj := &graph.DBGraph{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// SaveNode godoc
// @Summary saves or updates node
// @Id SaveNode
// @Accept  json
// @Produce  json
// @Param request body graph.DBNode true "request"
// @Success 200 {object} graph.DBNode
// @Router /node [post]
func (c *Controller) SaveNode(g *gin.Context) {
	c.DoCall(g, c.storage.SaveNode, func() (interface{}, error) {
		obj := &graph.DBNode{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// SaveConnection godoc
// @Summary saves or updates Connection
// @Id SaveConnection
// @Accept  json
// @Produce  json
// @Param request body graph.DBConnection true "request"
// @Success 200 {object} graph.DBConnection
// @Router /connection [post]
func (c *Controller) SaveConnection(g *gin.Context) {
	c.DoCall(g, c.storage.SaveConnection, func() (interface{}, error) {
		obj := &graph.DBConnection{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// SaveEventCard godoc
// @Summary saves or updates EventCard
// @Id SaveEventCard
// @Accept  json
// @Produce  json
// @Param request body graph.DBEventCard true "request"
// @Success 200 {object} graph.DBEventCard
// @Router /event-card [post]
func (c *Controller) SaveEventCard(g *gin.Context) {
	c.DoCall(g, c.storage.SaveEventCard, func() (interface{}, error) {
		obj := &graph.DBEventCard{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// DeleteGraph godoc
// @Summary deletes graph and all related Nodes, Connections and EventCards
// @Id DeleteGraph
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /graph [delete]
func (c *Controller) DeleteGraph(g *gin.Context) {
	c.DoCall(g, c.storage.DeleteGraph, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// DeleteNode godoc
// @Summary deletes Node
// @Id DeleteNode
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /node [delete]
func (c *Controller) DeleteNode(g *gin.Context) {
	c.DoCall(g, c.storage.DeleteNode, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// DeleteConnection godoc
// @Summary deletes Connection
// @Id DeleteConnection
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /connection [delete]
func (c *Controller) DeleteConnection(g *gin.Context) {
	c.DoCall(g, c.storage.DeleteConnection, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// DeleteEventCard godoc
// @Summary deletes EventCard
// @Id DeleteEventCard
// @Accept  json
// @Produce  json
// @Param request body storage.DeleteRequest true "request"
// @Success 200 {object} storage.DeleteResponse
// @Router /event-card [delete]
func (c *Controller) DeleteEventCard(g *gin.Context) {
	c.DoCall(g, c.storage.DeleteEventCard, func() (interface{}, error) {
		obj := &storage.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// ListModule godoc
// @Summary lists modules
// @Id ListModule
// @Accept  json
// @Produce  json
// @Success 200 {object} module.ListResponse
// @Router /list-module [get]
func (c *Controller) ListModule(g *gin.Context) {
	c.DoCall(g, c.module.List, func() (interface{}, error) {
		obj := &module.ListRequest{}
		err := g.BindQuery(obj)
		return obj, err
	})
}

// SaveModule godoc
// @Summary saves Module
// @Id SaveModule
// @Accept  json
// @Produce  json
// @Param request body module.DBModule true "request"
// @Success 200 {object} module.DBModule
// @Router /module [post]
func (c *Controller) SaveModule(g *gin.Context) {
	c.DoCall(g, c.module.Save, func() (interface{}, error) {
		obj := &module.DBModule{}
		err := g.BindJSON(obj)
		return obj, err
	})
}

// DeleteModule godoc
// @Summary deletes Module
// @Id DeleteModule
// @Accept  json
// @Produce  json
// @Param request body module.DeleteRequest true "request"
// @Success 200 {object} module.DeleteResponse
// @Router /module [delete]
func (c *Controller) DeleteModule(g *gin.Context) {
	c.DoCall(g, c.module.Delete, func() (interface{}, error) {
		obj := &module.DeleteRequest{}
		err := g.BindJSON(obj)
		return obj, err
	})
}
