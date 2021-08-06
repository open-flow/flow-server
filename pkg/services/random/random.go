package random

import (
	"gorm.io/gorm"
)

type Service struct {
	orm *gorm.DB
}

func New(orm *gorm.DB) *Service {
	return &Service{
		orm: orm,
	}
}

func (s *Service) StoreRandomGraph() error {
	//g := graph.Random()
	//
	//err := s.orm.
	//	Session(&gorm.Session{FullSaveAssociations: true}).
	//	Transaction(func(tx *gorm.DB) error {
	//		res := tx.Create(&g)
	//		return res.Error
	//	})
	//
	//if err != nil {
	//	return err
	//}

	return nil
}
