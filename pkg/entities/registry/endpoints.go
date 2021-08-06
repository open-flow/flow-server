package registry

type ModuleDef struct {
	Module    string
	Functions []*FunctionDef
}

type FunctionDef struct {
	Function string
}
