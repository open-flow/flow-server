package flow

import (
	"context"
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/flow-proto/gen/go/flow/v1"
	"gorm.io/gorm"
)

type batchService struct {
	db *gorm.DB
	api.UnimplementedBatchServiceServer
}

func NewBatchService(db *gorm.DB) api.BatchServiceServer {
	return &batchService{
		db: db,
	}
}

func (s *batchService) Save(ctx context.Context, request *api.BatchSaveRequest) (*api.BatchSaveResponse, error) {
	response := &api.BatchSaveResponse{
		ProjectId: request.ProjectId,
		GraphId:   request.GraphId,
	}

	graph := &Graph{}
	nodes := &[]*Node{}
	cards := &[]*EventCard{}
	connections := &[]*Connection{}

	err := copier.Copy(nodes, request.Nodes)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(cards, request.Cards)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(connections, request.Connections)
	if err != nil {
		return nil, err
	}

	for _, v := range *nodes {
		v.ProjectId = request.ProjectId
		v.GraphId = request.GraphId
		v.Id = 0
	}

	for _, v := range *cards {
		v.ProjectId = request.ProjectId
		v.GraphId = request.GraphId
		v.Id = 0
	}

	for _, v := range *connections {
		v.ProjectId = request.ProjectId
		v.GraphId = request.GraphId
		v.Id = 0
	}

	err = s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_Id = ? and Id = ?", request.ProjectId, request.GraphId).
				First(graph)

			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(nodes)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(cards)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(connections)
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	err = copier.Copy(&response.Nodes, nodes)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&response.Cards, cards)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&response.Connections, connections)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *batchService) Delete(ctx context.Context, request *api.BatchDeleteRequest) (*api.BatchDeleteResponse, error) {
	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"project_Id = ? and graph_Id = ? and Id in ?",
				request.ProjectId,
				request.GraphId,
				request.Connections,
			).Delete(&Connection{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_Id = ? and graph_Id = ? and Id in ?",
					request.ProjectId,
					request.GraphId,
					request.Cards,
				).Delete(&EventCard{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_Id = ? and graph_Id = ? and Id in ?",
					request.ProjectId,
					request.GraphId,
					request.Nodes,
				).Delete(&Node{})
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &api.BatchDeleteResponse{}, nil
}
