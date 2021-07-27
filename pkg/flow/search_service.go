package flow

import (
	"context"
	"github.com/jinzhu/copier"
	api "gitlab.com/yautoflow/flow-proto/gen/go/flow/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type searchService struct {
	db *gorm.DB
	api.UnimplementedSearchServiceServer
}

func NewSearchService(db *gorm.DB) api.SearchServiceServer {
	return &searchService{
		db: db,
	}
}

func (s *searchService) FindEventNode(ctx context.Context, req *api.FindNodeRequest) (*api.FindNodeResponse, error) {
	if req.Search.OwnerType == "" || req.Search.OwnerId == "" {
		return nil, status.Error(codes.NotFound, "owner_type and owner_id are mandatory")
	}

	res := &api.FindNodeResponse{}

	search := &EventCard{}
	err := copier.Copy(search, req.Search)
	if err != nil {
		return nil, err
	}

	var cards []*EventCard
	var graphs []*Graph

	err = s.db.Session(
		&gorm.Session{Context: ctx},
	).Transaction(func(tx *gorm.DB) error {
		res := tx.Where(search).Find(&cards)
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

	graphMap := make(map[uint64]*Graph)
	for _, g := range graphs {
		graphMap[g.Id] = g
	}

	cardsMap := make(map[uint64][]*EventCard)
	for _, c := range cards {
		slice, _ := cardsMap[c.GraphId]
		cardsMap[c.GraphId] = append(slice, c)
	}

	for _, g := range graphMap {
		eventGraph := api.EventGraph{
			Graph: &api.Graph{},
		}

		err = copier.CopyWithOption(eventGraph.Graph, g, DEEP_COPY)
		if err != nil {
			return nil, err
		}

		err = copier.CopyWithOption(&eventGraph.Nodes, g.Nodes, DEEP_COPY)
		if err != nil {
			return nil, err
		}

		err = copier.CopyWithOption(&eventGraph.Connections, g.Connections, DEEP_COPY)
		if err != nil {
			return nil, err
		}

		err = copier.CopyWithOption(&eventGraph.Cards, g.Cards, DEEP_COPY)
		if err != nil {
			return nil, err
		}

		err = copier.CopyWithOption(&eventGraph.ActiveCards, cardsMap[g.Id], DEEP_COPY)
		if err != nil {
			return nil, err
		}

		res.Graphs = append(res.Graphs, &eventGraph)
	}

	return res, nil
}
