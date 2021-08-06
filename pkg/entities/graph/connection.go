package graph

type DBConnection struct {
	IDGraph
	DataUI
	DataConnection
}

type DataConnection struct {
	//Placement
	LocalId uint64

	SourcePort string
	SourceId   uint64

	TargetPort string
	TargetId   uint64
}
