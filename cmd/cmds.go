package cmd

import (
	"autoflow/pkg/infra"
	"autoflow/pkg/services/batch"
	"autoflow/pkg/services/execution"
	"autoflow/pkg/services/random"
	"autoflow/pkg/services/registry"
	"autoflow/pkg/services/registry/static"
	"autoflow/pkg/services/schedule"
	"autoflow/pkg/services/search"
	"autoflow/pkg/services/storage"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var randomGraph = &cobra.Command{
	Use:   "random-graph",
	Short: "Generate and store random graph",
	Run: func(cmd *cobra.Command, args []string) {
		var random *random.Service
		fx.New(
			Provide(),
			fx.Populate(&random),
		)
		err := random.StoreRandomGraph()
		if err != nil {
			panic(err)
		}
	},
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Serve service",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			Provide(),
			fx.Invoke(infra.NewGin, static.HttpEndpointStaticConfig),
		)
		app.Run()
	},
}

func Provide() fx.Option {
	return fx.Provide(
		batch.New,
		storage.New,
		random.New,
		search.New,
		schedule.New,
		execution.NewExecuteService,
		infra.NewGin,
		infra.NewConfig,
		infra.NewGorm,
		static.HttpEndpointStaticConfig,
		registry.NewRegistryService,
		infra.NewLogger,
		infra.NewSugaredLogger,
	)
}
