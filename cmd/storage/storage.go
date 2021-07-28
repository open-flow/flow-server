package storage

import (
	"autoflow/inst"
	"autoflow/pkg/storage"
	"autoflow/pkg/storage/random"
	"github.com/spf13/cobra"
)

var Storage = &cobra.Command{
	Use:   "storage",
	Short: "Storage operations",
}

var randomGraph = &cobra.Command{
	Use:   "random-graph",
	Short: "Generate and store random graph",
	Run: func(cmd *cobra.Command, args []string) {
		db := inst.Gorm(
			inst.EnvGormConfig(),
		)
		service := random.NewService(db)
		err := service.StoreRandomGraph()

		if err != nil {
			panic(err)
		}
	},
}

var listen = &cobra.Command{
	Use:   "listen",
	Short: "Bind nats listeners",
	Run: func(cmd *cobra.Command, args []string) {
		db := inst.Gorm(
			inst.EnvGormConfig(),
		)
		nc := inst.Nats(
			inst.EnvNatsConfig(),
		)

		err := storage.ListenStorage(db, nc)

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	Storage.AddCommand(randomGraph)
	Storage.AddCommand(listen)
}
