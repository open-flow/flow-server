package infra

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

type FlowConfig struct {
	DSN           string
	HttpAddr      string
	HostName      string
	NatsUrl       string
	Development   bool
	ShowSql       bool
	RedisUsername string
	RedisPassword string
	RedisHost     string
	RedisPort     string
}

type RedisConfig struct {
	Addr map[string]string
}

func NewConfig() (*FlowConfig, error) {
	viper.SetDefault("DSN", "host=localhost user=postgres password=postgres dbname=flow port=5432 sslmode=disable TimeZone=Europe/Moscow")
	viper.SetDefault("HostName", "localhost:8080")
	viper.SetDefault("HttpAddr", ":8080")

	viper.SetDefault("RedisHost", "localhost")
	viper.SetDefault("RedisPort", "6379")
	viper.SetDefault("Development", false)
	viper.SetDefault("NatsUrl", "")
	viper.SetDefault("ShowSql", false)

	var config FlowConfig
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(&config)
	log.Print(string(bytes))
	return &config, nil
}
