package flow

import (
	"github.com/jinzhu/copier"
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
	response := &api.SaveResponse{
		ProjectID: request.ProjectID,
		GraphID:   request.GraphID,
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
		v.ProjectID = request.ProjectID
		v.GraphID = request.GraphID
		v.ID = 0
	}

	for _, v := range *cards {
		v.ProjectID = request.ProjectID
		v.GraphID = request.GraphID
		v.ID = 0
	}

	for _, v := range *connections {
		v.ProjectID = request.ProjectID
		v.GraphID = request.GraphID
		v.ID = 0
	}

	err = s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", request.ProjectID, request.GraphID).
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

func (s *batchService) Delete(ctx context.Context, request *api.DeleteRequest) (*api.DeleteResponse, error) {
	panic("implement me")
}
