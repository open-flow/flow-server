package infra

import (
	"autoflow/pkg/entities/graph"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(config *FlowConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.MySqlDSN), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&graph.DBGraph{}, &graph.DBEventCard{}, &graph.DBNode{}, &graph.DBConnection{})
	if err != nil {
		return nil, err
	}

	return db, err
}
