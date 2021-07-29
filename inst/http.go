package inst

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GinConfig struct {
	HttpAddr string
}

func NewGin(config *GinConfig) *gin.Engine {
	return gin.New()
}

func EnvGinConfig() *GinConfig {
	var config GinConfig
	viper.SetDefault("HttpAddr", ":8080")
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
