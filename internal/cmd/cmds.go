package cmd

import (
	"autoflow/internal/http"
	"autoflow/internal/infra"
	"autoflow/internal/modules/scallback"
	"autoflow/internal/modules/sendpoint"
	storage2 "autoflow/internal/modules/serrors"
	"autoflow/internal/modules/sgraph"
	"autoflow/internal/modules/srepo"
	"autoflow/internal/modules/sscheduler"
	"autoflow/pkg/common"
	"autoflow/pkg/engine/call"
	"autoflow/pkg/engine/state"
	"autoflow/pkg/storage/batch"
	"autoflow/pkg/storage/endpoint"
	"autoflow/pkg/storage/graph"
	"autoflow/pkg/storage/search"
	"autoflow/pkg/storage/storage"
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
				Add(batch.SaveRequest{}).
				Add(batch.SaveResponse{}).
				Add(batch.DeleteRequest{}).
				Add(endpoint.Endpoint{}).
				Add(endpoint.DBEndpoint{}).
				Add(endpoint.DataEndpoint{}).
				Add(endpoint.DBError{}).
				Add(endpoint.DataError{}).
				Add(endpoint.Container{}).
				Add(common.ProjectModel{}).
				Add(common.ProjectSpace{}).
				Add(call.Action{}).
				Add(call.Error{}).
				Add(call.Return{}).
				Add(call.CallbackRequest{}).
				Add(call.CallbackResponse{}).
				Add(state.State{}).
				Add(state.Cursor{}).
				Add(state.Memory{}).
				Add(graph.DBGraph{}).
				Add(graph.DBNode{}).
				Add(graph.DBEventCard{}).
				Add(graph.DBConnection{}).
				Add(graph.DataConnection{}).
				Add(graph.DataEvent{}).
				Add(graph.DataEventCard{}).
				Add(graph.DataNode{}).
				Add(graph.DataGraph{}).
				Add(search.FindActiveRequest{}).
				Add(search.FindActiveResponse{}).
				Add(search.ActiveGraph{}).
				Add(storage.ListGraphRequest{}).
				Add(storage.ListGraphResponse{})

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

		scallback.NewCallback,

		sendpoint.NewEndpointCache,
		sendpoint.NewEndpoint,
		storage2.NewErrors,

		storage2.NewLogger,

		sendpoint.NewRegistry,

		srepo.NewRepo,

		sgraph.NewGraphBatch,
		sgraph.NewGraph,
		sgraph.NewActive,
		sscheduler.NewSchedule,

		http.NewController,
	)
}
