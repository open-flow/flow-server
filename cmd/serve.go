package cmd

import (
	"autoflow/init"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve.go",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		db := init.Gorm(
			init.EnvGormConfig(),
		)

		init.StartGrpc(
			db,
			init.EnvGrpcConfig(),
		)
	},
}
