package search

import (
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/search"
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) FindActive(ctx context.Context, req *search.FindActiveRequest) (*search.FindActiveResponse, error) {
	if req.OwnerType == "" || req.OwnerId == "" {
		return nil, fmt.Errorf("owner_type and owner_id are mandatory")
	}

	var cards []*graph.DBEventCard
	var graphs []*graph.DBGraph

	err := s.db.Session(
		&gorm.Session{Context: ctx},
	).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&graph.DBEventCard{}).
			Where(req).
			Find(&cards)

		if res.Error != nil {
			return res.Error
		}

		var graphIds = make([]uint, len(cards))
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

	graphMap := make(map[uint]*graph.DBGraph)
	for _, g := range graphs {
		graphMap[g.Id] = g
	}

	cardsMap := make(map[uint][]*graph.DBEventCard)
	for _, c := range cards {
		slice, _ := cardsMap[c.GraphId]
		cardsMap[c.GraphId] = append(slice, c)
	}

	var activeGraphs []*search.ActiveGraph
	for _, g := range graphMap {
		active := &search.ActiveGraph{
			Graph:  g,
			Active: cardsMap[g.Id],
		}
		activeGraphs = append(activeGraphs, active)
	}

	return &search.FindActiveResponse{
		Graphs: activeGraphs,
	}, nil
}
