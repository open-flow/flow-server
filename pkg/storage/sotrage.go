package storage

import (
	"autoflow/pkg/dtos"
	"autoflow/pkg/orm"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StorageService struct {
	db *gorm.DB
}

func NewStorageService(db *gorm.DB) *StorageService {
	return &StorageService{
		db: db,
	}
}

func (s *StorageService) SaveGraph(c context.Context, graph *orm.Graph) (*orm.Graph, error) {
	err := s.StoreGeneric(c, graph)
	if err != nil {
		return nil, err
	}
	return graph, nil
}

func (s *StorageService) SaveNode(c context.Context, node *orm.Node) (*orm.Node, error) {
	err := s.StoreGeneric(c, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *StorageService) SaveEventCard(c context.Context, card *orm.EventCard) (*orm.EventCard, error) {
	err := s.StoreGeneric(c, card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (s *StorageService) SaveConnection(c context.Context, connection *orm.Connection) (*orm.Connection, error) {
	err := s.StoreGeneric(c, connection)
	if err != nil {
		return nil, err
	}
	return connection, nil
}

func (s *StorageService) DeleteGraph(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Graph{})
}

func (s *StorageService) DeleteNode(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Node{})
}

func (s *StorageService) DeleteEventCard(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.EventCard{})
}

func (s *StorageService) DeleteConnection(c context.Context, request *dtos.DeleteRequest) (*dtos.DeleteResponse, error) {
	return s.DeleteGeneric(c, request, &orm.Connection{})
}

func (s *StorageService) GetFullGraph(c context.Context, r *dtos.GetFullGraphRequest) (*orm.Graph, error) {
	var graph = &orm.Graph{}

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", r.ProjectId, r.Id).
				Preload(clause.Associations).
				First(graph)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return graph, nil
}

func (s *StorageService) ListGraph(c context.Context, r *dtos.ListGraphRequest) (*dtos.ListGraphResponse, error) {
	var graphs []*orm.Graph

	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id in ?", r.ProjectIds).
				Find(&graphs)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &dtos.ListGraphResponse{
		Graphs: graphs,
	}, nil
}

func (s *StorageService) StoreGeneric(c context.Context, model interface{}) error {
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			//todo change to use projectID
			res := tx.FirstOrCreate(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorageService) DeleteGeneric(c context.Context, req *dtos.DeleteRequest, model interface{}) (*dtos.DeleteResponse, error) {
	err := s.db.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", req.ProjectId, req.Id).
				Delete(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &dtos.DeleteResponse{}, nil
}
