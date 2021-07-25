package flow

import (
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type batchService struct {
	db *gorm.DB
}

func NewBatchService(db *gorm.DB) api.BatchServiceServer {
	return &batchService{
		db: db,
	}
}

func (s *batchService) Save(ctx context.Context, request *api.SaveRequest) (*api.SaveResponse, error) {
	panic("implement me")
}

func (s *batchService) Delete(ctx context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	panic("implement me")
}
