package graph

type DBConnection struct {
	IDGraph
	DataUI
	DataConnection
}

type DataConnection struct {
	//Placement
	LocalId uint

	SourcePort string
	SourceId   uint

	TargetPort string
	TargetId   uint
}
