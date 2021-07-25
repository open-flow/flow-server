package flow

import (
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Server struct {
	orm *gorm.DB
}

func NewServer(orm *gorm.DB) api.GraphServiceServer {
	return &Server{
		orm: orm,
	}
}

func (s *Server) StoreGraph(c context.Context, graph *api.Graph) (*api.Graph, error) {
	generic, err := s.StoreGeneric(c, graph)
	if err != nil {
		return nil, err
	}
	return generic.(*api.Graph), nil
}

func (s *Server) StoreNode(c context.Context, node *api.Node) (*api.Node, error) {
	generic, err := s.StoreGeneric(c, node)
	if err != nil {
		return nil, err
	}
	return generic.(*api.Node), nil
}

func (s *Server) StoreEvent(c context.Context, event *api.Event) (*api.Event, error) {
	generic, err := s.StoreGeneric(c, event)
	if err != nil {
		return nil, err
	}
	return generic.(*api.Event), nil
}

func (s *Server) StoreConnection(c context.Context, connection *api.Connection) (*api.Connection, error) {
	generic, err := s.StoreGeneric(c, connection)
	if err != nil {
		return nil, err
	}
	return generic.(*api.Connection), nil
}

func (s *Server) StoreBatch(c context.Context, request *api.BatchRequest) (*api.BatchResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) RemoveBatch(c context.Context, remove *api.EdgedBatchRemove) (*api.RemoveResponse, error) {
	err := s.orm.
		Session(&gorm.Session{
			Context: c,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"ProjectID = ? AND GraphID = ? AND ID in ?",
				remove.ProjectID,
				remove.GraphID,
				remove.Connections,
			).Delete(&Connection{})

			if res.Error != nil {
				return res.Error
			}

			res = tx.Where(
				"ProjectID = ? AND GraphID = ? AND ID in ?",
				remove.ProjectID,
				remove.GraphID,
				remove.Events,
			).Delete(&Event{})

			if res.Error != nil {
				return res.Error
			}

			res = tx.Where(
				"ProjectID = ? AND GraphID = ? AND ID in ?",
				remove.ProjectID,
				remove.GraphID,
				remove.Nodes,
			).Delete(&Node{})

			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &api.RemoveResponse{
		Ok: true,
	}, nil
}

func (s *Server) RemoveGraph(c context.Context, request *api.GraphRemoveRequest) (*api.RemoveResponse, error) {
	err := s.orm.
		Session(&gorm.Session{
			Context: c,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"ProjectID = ? AND ID = ?",
				request.ProjectID, request.GraphID,
			).Delete(&Graph{})

			if res.Error != nil {
				return res.Error
			}

			return nil
		})
	if err != nil {
		return nil, err
	}
	return &api.RemoveResponse{
		Ok: true,
	}, nil
}

func (s *Server) GetGraph(c context.Context, request *api.GetGraphRequest) (*api.GetGraphResponse, error) {
	db := s.orm.Session(&gorm.Session{
		Context: c,
	})
	var entity Graph
	var apiGraph api.Graph
	res := db.
		Where("project_id = ? and id = ?", request.ProjectID, request.GraphID).
		Preload(clause.Associations).
		Preload("Events.Cards").
		First(&entity)
	if res.Error != nil {
		return nil, res.Error
	}
	err := copier.CopyWithOption(&apiGraph, &entity, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}
	return &api.GetGraphResponse{
		Graph: &apiGraph,
	}, nil
}

func (s *Server) ShallowGraphList(c context.Context, request *api.ListGraphsRequest) (*api.ListGraphResponse, error) {
	db := s.orm.Session(&gorm.Session{
		Context: c,
	})
	var graphList []*Graph
	var apiGraphList []*api.Graph

	db.Where("project_id in ?", request.ProjectID).Find(&graphList)
	err := copier.CopyWithOption(&apiGraphList, graphList, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	return &api.ListGraphResponse{
		Graph: apiGraphList,
	}, nil
}
