package flow

import (
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	var entity Graph
	entity.assignID(graph)

	if graph.ID != 0 {
		res := s.orm.First(&entity)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		res := s.orm.Create(&entity)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	if entity.assign(graph) {
		res := s.orm.Save(&entity)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) StoreNode(c context.Context, node *api.Node) (*api.Node, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) StoreEvent(c context.Context, node *api.Node) (*api.Event, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *Server) StoreBatch(c context.Context, request *api.BatchRequest) (*api.BatchResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
