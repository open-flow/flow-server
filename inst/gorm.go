package inst

import (
	"autoflow/pkg/flow"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
import "gorm.io/driver/mysql"

var Gorm *gorm.DB

func InitGorm() {
	var err error

	Gorm, err = gorm.Open(mysql.Open(Config.MySqlDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	err = Gorm.AutoMigrate(&flow.Graph{}, &flow.Event{}, &flow.EventCard{}, &flow.Node{})
	if err != nil {
		panic(err)
	}
}
