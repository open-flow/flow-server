package batch

import (
	"autoflow/pkg/entities/batch"
	"autoflow/pkg/entities/graph"
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
	g := &graph.DBGraph{}

	cards := make([]*graph.DBEventCard, len(r.Cards))
	nodes := make([]*graph.DBNode, len(r.Nodes))
	connections := make([]*graph.DBConnection, len(r.Connections))

	idGraph := graph.IDGraph{
		IDProject: graph.IDProject{
			ProjectID: r.ProjectID,
		},
		GraphID: r.ID,
	}

	for i, c := range r.Cards {
		cards[i] = &graph.DBEventCard{
			IDGraph:       idGraph,
			DataEventCard: c.DataEventCard,
			DataUI:        c.DataUI,
		}
	}

	for i, c := range r.Nodes {
		nodes[i] = &graph.DBNode{
			IDGraph:  idGraph,
			DataNode: c.DataNode,
			DataUI:   c.DataUI,
		}
	}

	for i, c := range r.Connections {
		connections[i] = &graph.DBConnection{
			IDGraph:        idGraph,
			DataConnection: c.DataConnection,
			DataUI:         c.DataUI,
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
				First(g)

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
		IDProject:   r.IDProject,
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
			).Delete(&graph.DBConnection{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectID,
					r.ID,
					r.Cards,
				).Delete(&graph.DBEventCard{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectID,
					r.ID,
					r.Nodes,
				).Delete(&graph.DBNode{})
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
