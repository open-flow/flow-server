package inst

import (
	"autoflow/pkg/flow"
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitGrpc() {
	lis, err := net.Listen("tcp", Config.HostAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGraphServiceServer(s, flow.NewServer(Gorm))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
