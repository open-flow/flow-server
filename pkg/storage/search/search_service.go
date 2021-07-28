package search

import (
	"context"
	"fmt"
	"gitlab.com/yautoflow/interfaces/dtos"
	"gitlab.com/yautoflow/interfaces/orm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) FindEventNode(ctx context.Context, req dtos.EventSearchRequest) (*dtos.EventSearchResponse, error) {
	if req.OwnerType == "" || req.OwnerId == "" {
		return nil, fmt.Errorf("owner_type and owner_id are mandatory")
	}

	var cards []*orm.EventCard
	var graphs []*orm.Graph

	err := s.db.Session(
		&gorm.Session{Context: ctx},
	).Transaction(func(tx *gorm.DB) error {
		res := tx.
			Where(req).
			Find(&cards)
		if res.Error != nil {
			return res.Error
		}

		var graphIds = make([]uint64, len(cards))
		for i, card := range cards {
			graphIds[i] = card.GraphId
		}
		res = tx.Where("id in ?", graphIds).
			Preload(clause.Associations).
			First(&graphs)
		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	graphMap := make(map[uint64]*orm.Graph)
	for _, g := range graphs {
		graphMap[g.Id] = g
	}

	cardsMap := make(map[uint64][]*orm.EventCard)
	for _, c := range cards {
		slice, _ := cardsMap[c.GraphId]
		cardsMap[c.GraphId] = append(slice, c)
	}

	var activeGraphs []*dtos.ActiveGraph
	for _, g := range graphMap {
		active := &dtos.ActiveGraph{
			Graph:       g,
			ActiveCards: cardsMap[g.Id],
		}
		activeGraphs = append(activeGraphs, active)
	}

	return &dtos.EventSearchResponse{
		Graphs: activeGraphs,
	}, nil
}
