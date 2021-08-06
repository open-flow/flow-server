package graph

type DBNode struct {
	IDGraph
	DataUI
	DataNode
}

type DataNode struct {
	//Placement
	LocalId uint64

	//Invocation
	Type      string
	Module    string
	Function  string
	Arguments string
}
