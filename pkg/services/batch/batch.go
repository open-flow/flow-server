package batch

import (
	"autoflow/pkg/entities/batch"
	graph2 "autoflow/pkg/entities/graph"
	"context"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Save(ctx context.Context, r *batch.SaveRequest) (*batch.SaveResponse, error) {
	graph := &graph2.DBGraph{}

	cards := make([]*graph2.DBEventCard, len(r.Cards))
	nodes := make([]*graph2.DBNode, len(r.Nodes))
	connections := make([]*graph2.DBConnection, len(r.Connections))

	for i, c := range r.Cards {
		cards[i] = &graph2.DBEventCard{
			DataEventCard: c,
		}
	}

	for i, c := range r.Nodes {
		nodes[i] = &graph2.DBNode{
			DataNode: c,
		}
	}

	for i, c := range r.Connections {
		connections[i] = &graph2.DBConnection{
			DataConnection: c,
		}
	}

	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.
				Where("project_id = ? and id = ?", r.ProjectID, r.ID).
				First(graph)

			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(cards)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(nodes)
			if res.Error != nil {
				return res.Error
			}

			res = tx.Create(connections)
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &batch.SaveResponse{
		IDGraph:     r.IDGraph,
		Nodes:       nodes,
		Cards:       cards,
		Connections: connections,
	}, nil
}

func (s *Service) Delete(ctx context.Context, r *batch.DeleteRequest) (*batch.DeleteResponse, error) {
	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"project_id = ? and graph_id = ? and id in ?",
				r.ProjectID,
				r.ID,
				r.Connections,
			).Delete(&graph2.DBConnection{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectID,
					r.ID,
					r.Cards,
				).Delete(&graph2.DBEventCard{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectID,
					r.ID,
					r.Nodes,
				).Delete(&graph2.DBNode{})
			if res.Error != nil {
				return res.Error
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	return &batch.DeleteResponse{}, nil
}
