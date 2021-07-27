package inst

import (
	"autoflow/pkg/flow"
	"github.com/spf13/viper"
	api "gitlab.com/yautoflow/flow-proto/gen/go/flow/v1"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
)

type GrpcConfig struct {
	HostAddr string
}

func StartGrpc(db *gorm.DB, config *GrpcConfig) {
	lis, err := net.Listen("tcp", config.HostAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterGraphServiceServer(s, flow.NewGraphService(db))
	api.RegisterBatchServiceServer(s, flow.NewBatchService(db))
	api.RegisterSearchServiceServer(s, flow.NewSearchService(db))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve.go: %v", err)
	}
}

func EnvGrpcConfig() *GrpcConfig {
	viper.SetDefault("HostAddr", ":9090")
	var config GrpcConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
