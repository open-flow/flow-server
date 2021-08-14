package main

import (
	"autoflow/internal/cmd"
)

//go:generate swagger generate spec -c pkg.engine.* -m -i api/endpoint/source.yaml -o api/endpoint/swagger.yaml

func main() {
	cmd.Execute()
}
