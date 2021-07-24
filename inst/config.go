package inst

import "github.com/spf13/viper"

type ConfigStruct struct {
	Development bool
	MySqlDSN    string
	HostAddr    string
}

type MongoConfig struct {
}

var Config ConfigStruct

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("Development", true)
	viper.SetDefault("MySqlDSN", "root:mysql@tcp(127.0.0.1:3306)/flow?charset=utf8mb4")
	viper.SetDefault("HostAddr", ":9090")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Config)
}
