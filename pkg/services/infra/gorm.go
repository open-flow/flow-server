package infra

import (
	"autoflow/pkg/orm"
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

	err = db.AutoMigrate(&orm.Graph{}, &orm.EventCard{}, &orm.Node{}, &orm.Connection{})
	if err != nil {
		return nil, err
	}

	return db, err
}
