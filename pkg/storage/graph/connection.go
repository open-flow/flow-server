package graph

import "autoflow/pkg/common"

type DBConnection struct {
	GraphObject
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
