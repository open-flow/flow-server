package flow

import (
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

var _ api.GraphServiceServer = &Server{}

func (s *Server) StoreGraph(c context.Context, graph *api.Graph) (*api.Graph, error) {
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
