package main

import (
	"autoflow/inst"
)

func main() {
	inst.InitConfig()
	inst.InitLogger()
	inst.InitGorm()
	inst.InitGrpc()
}
