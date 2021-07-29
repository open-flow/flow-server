package inst

import (
	"autoflow/pkg/orm"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormConfig struct {
	MySqlDSN string
}

func Gorm(config *GormConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.MySqlDSN), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&orm.Graph{}, &orm.EventCard{}, &orm.Node{}, &orm.Connection{})
	if err != nil {
		panic(err)
	}

	return db
}

func EnvGormConfig() *GormConfig {
	var config GormConfig
	viper.SetDefault("MySqlDSN", "root:mysql@tcp(127.0.0.1:3306)/flow?charset=utf8mb4")
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
