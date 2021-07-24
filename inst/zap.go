package inst

import "go.uber.org/zap"

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

func InitLogger() {
	var err error
	if Config.Development {
		Logger, err = zap.NewDevelopment()
	} else {
		Logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(err)
	}
	Sugar = Logger.Sugar()
}
