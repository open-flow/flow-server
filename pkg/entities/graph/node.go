package graph

type DBNode struct {
	IDGraph
	DataUI
	DataNode
}

type DataNode struct {
	//Placement
	LocalId uint `json:"localId"`

	//Invocation
	Type      string `json:"type"`
	Module    string `json:"module"`
	Function  string `json:"function"`
	Arguments string `json:"arguments"`
}
