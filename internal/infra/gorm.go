package infra

import (
	"autoflow/pkg/entities/graph"
	"autoflow/pkg/entities/module"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
)

func NewGorm(config *FlowConfig) (*gorm.DB, error) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NameReplacer: strings.NewReplacer("db_", ""),
		},
	}
	if config.ShowSql {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(config.MySqlDSN), gormConfig)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&graph.DBGraph{},
		&graph.DBEventCard{},
		&graph.DBNode{},
		&graph.DBConnection{},
		&module.DBModule{},
	)
	if err != nil {
		return nil, err
	}

	return db, err
}
