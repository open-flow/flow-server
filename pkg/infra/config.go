package infra

import (
	"github.com/spf13/viper"
)

type FlowConfig struct {
	MySqlDSN    string
	HttpAddr    string
	Development bool
	ShowSql     bool
}

func NewConfig() (*FlowConfig, error) {
	viper.SetDefault("MySqlDSN", "root:mysql@tcp(127.0.0.1:3306)/flow?charset=utf8mb4")
	viper.SetDefault("HttpAddr", ":8080")
	viper.SetDefault("Development", true)
	viper.SetDefault("ShowSql", false)

	var config FlowConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
