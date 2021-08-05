package registry

type EndpointDef struct {
	Module    string
	Functions []*EndpointFunctionDef
}

type EndpointFunctionDef struct {
	Function string
}
