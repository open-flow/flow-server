package storage

import (
	utils2 "autoflow/pkg/utils"
	"gorm.io/gorm"
)

type RandomService struct {
	orm *gorm.DB
}

func NewRandomService(orm *gorm.DB) *RandomService {
	return &RandomService{
		orm: orm,
	}
}

func (s *RandomService) StoreRandomGraph() error {
	var graph = utils2.Graph()

	err := s.orm.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Create(&graph)
			return res.Error
		})

	if err != nil {
		return err
	}

	return nil
}
