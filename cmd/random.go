package cmd

import (
	"autoflow/inst"
	"autoflow/pkg/flow/random"
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
