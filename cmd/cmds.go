package cmd

import (
	"autoflow/pkg/services/execution"
	infra2 "autoflow/pkg/services/infra"
	"autoflow/pkg/services/registry"
	"autoflow/pkg/services/registry/static"
	"autoflow/pkg/services/storage"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var randomGraph = &cobra.Command{
	Use:   "random-graph",
	Short: "Generate and store random graph",
	Run: func(cmd *cobra.Command, args []string) {
		var random *storage.RandomService
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
			fx.Invoke(infra2.NewGin, static.HttpEndpointStaticConfig),
		)
		app.Run()
	},
}

func Provide() fx.Option {
	return fx.Provide(
		storage.NewBatchService,
		storage.NewGraphService,
		storage.NewRandomService,
		storage.NewSearchService,
		execution.NewScheduleService,
		execution.NewExecuteService,
		infra2.NewGin,
		infra2.NewConfig,
		infra2.NewGorm,
		static.HttpEndpointStaticConfig,
		registry.NewRegistryService,
		infra2.NewLogger,
		infra2.NewSugaredLogger,
	)
}
