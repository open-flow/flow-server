package graph

type DBConnection struct {
	IDGraph
	DataUI
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
