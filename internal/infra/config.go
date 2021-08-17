package infra

import (
	"github.com/spf13/viper"
)

type FlowConfig struct {
	DSN         string
	HttpAddr    string
	HostName    string
	NatsUrl     string
	Development bool
	ShowSql     bool
	Redis       *RedisConfig
}

type RedisConfig struct {
	Addr map[string]string
}

func NewConfig() (*FlowConfig, error) {
	viper.SetDefault("DSN", "host=localhost user=postgres password=postgres dbname=flow port=5432 sslmode=disable TimeZone=Europe/Moscow")
	viper.SetDefault("HostName", "localhost:8080")
	viper.SetDefault("HttpAddr", ":8080")
	viper.SetDefault("Redis", &RedisConfig{
		Addr: map[string]string{
			"server1": ":6379",
		},
	})
	viper.SetDefault("Development", true)
	viper.SetDefault("NatsUrl", "")
	viper.SetDefault("ShowSql", false)

	var config FlowConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
