package cmd

import (
	"autoflow/inst"
	"autoflow/pkg/storage/binding"
	"autoflow/pkg/storage/random"
	"github.com/spf13/cobra"
)

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
	Use:   "serve",
	Short: "Serve service",
	Run: func(cmd *cobra.Command, args []string) {
		db := inst.Gorm(
			inst.EnvGormConfig(),
		)
		config := inst.EnvGinConfig()
		e := inst.NewGin(
			config,
		)
		binding.BindStorageGin(e, db)
		_ = e.Run(config.HttpAddr)
	},
}
