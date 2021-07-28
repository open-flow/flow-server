package random

import (
	"gorm.io/gorm"
)

type Service struct {
	orm *gorm.DB
}

func NewService(orm *gorm.DB) *Service {
	return &Service{
		orm: orm,
	}
}

func (s *Service) StoreRandomGraph() error {
	var graph = Graph()

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
