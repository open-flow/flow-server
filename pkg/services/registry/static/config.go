package static

import (
	"autoflow/pkg/services/registry"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type StaticHttpConfig struct {
	Endpoints []struct {
		Url    string
		Module string
	}
}

func HttpEndpointStaticConfig(ls fx.Lifecycle, svc *registry.RegistryService, logger *zap.Logger) (*StaticHttpConfig, error) {
	config := &StaticHttpConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	ls.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			for _, e := range config.Endpoints {
				endpoint := registry.NewHttpEndpoint(e.Url, e.Module, logger)
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
