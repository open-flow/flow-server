package graph

type DBGraph struct {
	IDGraph

	DataUI
	DataGraph

	Nodes       []DBNode       `gorm:"foreignKey:ID;references:GraphID;constraint:OnDelete:CASCADE;"`
	Cards       []DBEventCard  `gorm:"foreignKey:ID;references:GraphID;constraint:OnDelete:CASCADE;"`
	Connections []DBConnection `gorm:"foreignKey:ID;references:GraphID;constraint:OnDelete:CASCADE;"`
}

type DataGraph struct {
	Counter uint64
}
