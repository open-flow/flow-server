package static

import (
	registry2 "autoflow/pkg/services/registry"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HttpConfig struct {
	Endpoints []struct {
		Url    string
		Module string
	}
}

func HttpEndpointStaticConfig(ls fx.Lifecycle, svc *registry2.Service, logger *zap.Logger) (*HttpConfig, error) {
	config := &HttpConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, e := range config.Endpoints {
				endpoint := registry2.NewHttpEndpoint(e.Url, e.Module, logger)
				err := endpoint.Initialize()
				if err != nil {
					return err
				}
				svc.RegisterEndpoint(endpoint)
			}
			return nil
		},
	})

	return config, nil
}
