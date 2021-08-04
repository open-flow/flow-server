package cmd

import (
	inst2 "autoflow/pkg/infra"
	"autoflow/pkg/runner"
	"autoflow/pkg/storage"
	"github.com/gin-gonic/gin"
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
		var g *gin.Engine
		var config *inst2.FlowConfig
		fx.New(
			Provide(),
			fx.Populate(&g, &config),
		)

		err := g.Run(config.HttpAddr)
		if err != nil {
			panic(err)
		}
	},
}

func Provide() fx.Option {
	return fx.Provide(
		storage.NewBatchService,
		storage.NewStorageService,
		storage.NewRandomService,
		storage.NewSearchService,
		runner.NewScheduleService,
		runner.NewExecuteService,
		inst2.NewGin,
		inst2.NewConfig,
		inst2.NewGorm,
	)
}
