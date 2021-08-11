package cmd

import (
	"autoflow/internal/http"
	"autoflow/internal/infra"
	"autoflow/internal/modules/engine/callback"
	"autoflow/internal/modules/engine/registry"
	"autoflow/internal/modules/engine/schedule"
	"autoflow/internal/modules/storage/batch"
	"autoflow/internal/modules/storage/endpoint"
	"autoflow/internal/modules/storage/logger"
	"autoflow/internal/modules/storage/repo"
	"autoflow/internal/modules/storage/search"
	"autoflow/internal/modules/storage/storage"
	"autoflow/pkg/common"
	callDto "autoflow/pkg/engine/call"
	stateDto "autoflow/pkg/engine/state"
	batchDto "autoflow/pkg/storage/batch"
	endpointDto "autoflow/pkg/storage/endpoint"
	"autoflow/pkg/storage/graph"
	searchDto "autoflow/pkg/storage/search"
	storageDto "autoflow/pkg/storage/storage"
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
				Add(batchDto.SaveRequest{}).
				Add(batchDto.SaveResponse{}).
				Add(batchDto.DeleteRequest{}).
				Add(endpointDto.Endpoint{}).
				Add(endpointDto.DBEndpoint{}).
				Add(endpointDto.DataEndpoint{}).
				Add(endpointDto.DBError{}).
				Add(endpointDto.DataError{}).
				Add(endpointDto.Container{}).
				Add(common.IDProject{}).
				Add(common.ByProjectId{}).
				Add(callDto.Action{}).
				Add(callDto.Error{}).
				Add(callDto.Return{}).
				Add(callDto.Request{}).
				Add(callDto.Response{}).
				Add(stateDto.State{}).
				Add(stateDto.Cursor{}).
				Add(stateDto.Memory{}).
				Add(graph.DBGraph{}).
				Add(graph.DBNode{}).
				Add(graph.DBEventCard{}).
				Add(graph.DBConnection{}).
				Add(graph.DataConnection{}).
				Add(graph.DataEvent{}).
				Add(graph.DataEventCard{}).
				Add(graph.DataNode{}).
				Add(graph.DataGraph{}).
				Add(searchDto.FindActiveRequest{}).
				Add(searchDto.FindActiveResponse{}).
				Add(searchDto.ActiveGraph{}).
				Add(storageDto.ListGraphRequest{}).
				Add(storageDto.ListGraphResponse{})

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
