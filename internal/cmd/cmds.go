package cmd

import (
	"autoflow/internal/http"
	"autoflow/internal/infra"
	"autoflow/internal/services/batch"
	"autoflow/internal/services/callback"
	"autoflow/internal/services/endpoint"
	"autoflow/internal/services/logger"
	"autoflow/internal/services/registry"
	"autoflow/internal/services/repo"
	"autoflow/internal/services/schedule"
	"autoflow/internal/services/search"
	"autoflow/internal/services/storage"
	batchDto "autoflow/pkg/entities/batch"
	"autoflow/pkg/entities/graph"
	"github.com/spf13/cobra"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"go.uber.org/fx"
)

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Serve service",
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			Provide(),
			fx.Invoke(http.NewController),
		)
		app.Run()
	},
}

var ts = &cobra.Command{
	Use:   "ts",
	Short: "Generate typescript interfaces",
	Run: func(cmd *cobra.Command, args []string) {
		converter :=
			typescriptify.New().
				WithBackupDir("ts").WithInterface(true).
				Add(graph.DBGraph{}).
				Add(graph.DBNode{}).
				Add(graph.DBEventCard{}).
				Add(graph.DBConnection{}).
				Add(graph.DataConnection{}).
				Add(graph.DataEvent{}).
				Add(graph.DataEventCard{}).
				Add(graph.DataNode{}).
				Add(graph.DataGraph{}).
				Add(batchDto.SaveRequest{}).
				Add(batchDto.SaveResponse{}).
				Add(batchDto.DeleteRequest{}).
				Add(batchDto.DeleteResponse{})

		err := converter.ConvertToFile("ts/models.ts")
		if err != nil {
			panic(err.Error())
		}
	},
}

func Provide() fx.Option {
	return fx.Provide(
		//infra
		infra.NewConfig,
		infra.NewNats,
		infra.NewGorm,
		infra.NewLogger,
		infra.NewSugaredLogger,
		infra.NewRedis,

		callback.New,

		endpoint.NewCache,
		endpoint.NewController,
		endpoint.NewErrorService,

		logger.NewService,

		registry.NewResty,
		registry.NewService,

		repo.NewService,

		batch.New,
		storage.New,
		search.New,
		schedule.New,

		http.NewController,
	)
}
