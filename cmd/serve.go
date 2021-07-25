package cmd

import (
	"autoflow/inst"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		db := inst.Gorm(
			inst.EnvGormConfig(),
		)

		inst.StartGrpc(
			db,
			inst.EnvGrpcConfig(),
		)
	},
}
