package graph

import "autoflow/pkg/common"

type DBNode struct {
	IDGraph
	common.DataUI
	DataNode
}

type DataNode struct {
	//Placement
	LocalId uint `json:"localId"`

	//Invocation
	Module    string `json:"module"`
	Function  string `json:"function"`
	Arguments string `json:"arguments"`
}
