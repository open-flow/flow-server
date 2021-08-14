package main

import (
	"autoflow/internal/cmd"
)

//go:generate swagger generate spec -m -i docs/source.yaml -o docs/swagger.yaml

func main() {
	cmd.Execute()
}
