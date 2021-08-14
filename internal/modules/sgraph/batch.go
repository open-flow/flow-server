package sgraph

import (
	"autoflow/pkg/common"
	"autoflow/pkg/storage/batch"
	"autoflow/pkg/storage/graph"
	"context"
	"gorm.io/gorm"
)

type GraphBatch struct {
	db *gorm.DB
}

func NewGraphBatch(db *gorm.DB) *GraphBatch {
	return &GraphBatch{
		db: db,
	}
}

func (s *GraphBatch) Save(ctx context.Context, r *batch.SaveRequest) (*batch.SaveResponse, error) {
	g := &graph.DBGraph{}

	cards := make([]*graph.DBEventCard, len(r.Cards))
	nodes := make([]*graph.DBNode, len(r.Nodes))
	connections := make([]*graph.DBConnection, len(r.Connections))

	idGraph := graph.GraphObject{
		ProjectModel: common.ProjectModel{
			ProjectId: r.ProjectId,
		},
		GraphId: r.Id,
	}

	for i, c := range r.Cards {
		cards[i] = &graph.DBEventCard{
			GraphObject:   idGraph,
			DataEventCard: c.DataEventCard,
			DataUI:        c.DataUI,
		}
	}

	for i, c := range r.Nodes {
		nodes[i] = &graph.DBNode{
			GraphObject: idGraph,
			DataNode:    c.DataNode,
			DataUI:      c.DataUI,
		}
	}

	for i, c := range r.Connections {
		connections[i] = &graph.DBConnection{
			GraphObject:    idGraph,
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
				Where("project_id = ? and id = ?", r.ProjectId, r.Id).
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
		ProjectModel: r.ProjectModel,
		Nodes:        nodes,
		Cards:        cards,
		Connections:  connections,
	}, nil
}

func (s *GraphBatch) Delete(ctx context.Context, r *batch.DeleteRequest) error {
	err := s.db.
		Session(&gorm.Session{
			Context:         ctx,
			CreateBatchSize: 15,
		}).
		Transaction(func(tx *gorm.DB) error {
			res := tx.Where(
				"project_id = ? and graph_id = ? and id in ?",
				r.ProjectId,
				r.Id,
				r.Connections,
			).Delete(&graph.DBConnection{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectId,
					r.Id,
					r.Cards,
				).Delete(&graph.DBEventCard{})
			if res.Error != nil {
				return res.Error
			}

			res = tx.
				Where(
					"project_id = ? and graph_id = ? and id in ?",
					r.ProjectId,
					r.Id,
					r.Nodes,
				).Delete(&graph.DBNode{})
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
