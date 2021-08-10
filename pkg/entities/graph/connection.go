package graph

import "autoflow/pkg/entities/common"

type DBConnection struct {
	IDGraph
	common.DataUI
	DataConnection
}

type DataConnection struct {
	//Placement
	LocalID uint `json:"localId"`

	SourcePort string `json:"sourcePort"`
	SourceID   uint   `json:"sourceId"`

	TargetPort string `json:"targetPort"`
	TargetId   uint   `json:"targetId"`
}
