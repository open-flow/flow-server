package flow

import (
	"github.com/jinzhu/copier"
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
	var err error
	var entity Graph

	err = s.orm.Session(&gorm.Session{FullSaveAssociations: true}).Transaction(func(tx *gorm.DB) error {
		res := tx.FirstOrCreate(&entity, graph)
		if res.Error != nil {
			return res.Error
		}

		err := copier.CopyWithOption(&entity, graph, copier.Option{DeepCopy: true})
		if err != nil {
			return err
		}

		res = tx.Save(&entity)
		if res.Error != nil {
			return res.Error
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	var resultGraph api.Graph
	err = copier.CopyWithOption(&resultGraph, &entity, copier.Option{
		DeepCopy: true,
	})
	if err != nil {
		return nil, err
	}
	return &resultGraph, nil
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
