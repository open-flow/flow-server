package registry

type ModuleDef struct {
	Module    string         `json:"module"`
	Functions []*FunctionDef `json:"functions"`
}

type FunctionDef struct {
	Function string `json:"function"`
}
