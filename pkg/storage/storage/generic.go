package storage

import (
	"autoflow/pkg/storage/dtos"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (s *GraphService) StoreGeneric(c context.Context, model interface{}) error {
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

func (s *GraphService) DeleteGeneric(c context.Context, req *dtos.DeleteRequest, model interface{}) (*dtos.DeleteResponse, error) {
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
