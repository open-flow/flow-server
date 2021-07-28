package inst

import (
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

type NatsConfig struct {
	NatsURL string
}

func Nats(config *NatsConfig) *nats.Conn {
	nc, err := nats.Connect(config.NatsURL)

	if err != nil {
		panic(err)
	}

	return nc
}

func EnvNatsConfig() *NatsConfig {
	var config NatsConfig
	viper.SetDefault("NatsURL", nats.DefaultURL)
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
