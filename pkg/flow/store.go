package flow

import (
	"fmt"
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/protorepo-flow-server-go"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

func (s *Server) StoreGeneric(c context.Context, apiObj interface{}) (interface{}, error) {
	var entity interface{}

	switch apiObj.(type) {
	case *api.Graph:
		entity = &Graph{}
	case *api.Node:
		entity = &Node{}
	case *api.Connection:
		entity = &Connection{}
	case *api.Event:
		entity = &Event{}
	default:
		fmt.Print("unknown type passed")
	}

	err := copier.CopyWithOption(entity, apiObj, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	err = s.orm.
		Session(&gorm.Session{Context: c}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Save(entity)
			if res.Error != nil {
				return res.Error
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	err = copier.CopyWithOption(apiObj, entity, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, err
	}

	return apiObj, nil
}
