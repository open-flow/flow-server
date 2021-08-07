package cmd

import (
	"autoflow/internal/http"
	"autoflow/internal/infra"
	"autoflow/internal/services/batch"
	"autoflow/internal/services/callback"
	"autoflow/internal/services/module"
	"autoflow/internal/services/registry"
	"autoflow/internal/services/registry/static"
	"autoflow/internal/services/schedule"
	"autoflow/internal/services/search"
	"autoflow/internal/services/storage"
	batchDto "autoflow/pkg/entities/batch"
	executionDto "autoflow/pkg/entities/execution"
	"autoflow/pkg/entities/graph"
	registryDto "autoflow/pkg/entities/module"
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
			fx.Invoke(http.NewController, static.HttpEndpointStaticConfig),
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
				Add(batchDto.DeleteResponse{}).
				Add(registryDto.FunctionDef{}).
				Add(registryDto.ModuleDef{}).
				Add(executionDto.CallAction{}).
				Add(executionDto.CallAction{}).
				Add(executionDto.CallError{}).
				Add(executionDto.CallReturn{}).
				Add(executionDto.State{}).
				Add(executionDto.Cursor{}).
				Add(executionDto.Memory{}).
				Add(executionDto.Request{}).
				Add(executionDto.Response{}).
				Add(graph.IDGraph{})

		err := converter.ConvertToFile("ts/models.ts")
		if err != nil {
			panic(err.Error())
		}
	},
}

func Provide() fx.Option {
	return fx.Provide(
		batch.New,
		storage.New,
		search.New,
		schedule.New,
		callback.New,
		module.New,
		http.NewController,
		infra.NewConfig,
		infra.NewNats,
		infra.NewGorm,
		static.HttpEndpointStaticConfig,
		registry.NewRegistryService,
		infra.NewLogger,
		infra.NewSugaredLogger,
	)
}
