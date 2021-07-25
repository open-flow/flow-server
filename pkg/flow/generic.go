package flow

import (
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type GenericService interface {
	GetDB() *gorm.DB
}

type ApiObject interface {
	GetID() uint64
	GetProjectID() uint64
}

func StoreGeneric(s GenericService, c context.Context, obj ApiObject, model interface{}) error {
	err := s.GetDB().
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ? and id != 0", obj.GetProjectID(), obj.GetID()).
				Assign(obj).
				FirstOrCreate(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return err
	}

	err = copier.Copy(obj, model)
	if err != nil {
		return err
	}

	return nil
}

func DeleteGeneric(s GenericService, c context.Context, req *api.IDRequest, model interface{}) (*api.DeleteResponse, error) {
	err := s.GetDB().
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", req.GetProjectID(), req.GetID()).
				Delete(model)

			if res.Error != nil {
				return res.Error
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &api.DeleteResponse{
		Ok: true,
	}, nil
}
