package graph

type DBGraph struct {
	IDProject

	DataUI
	DataGraph

	Nodes       []DBNode       `gorm:"foreignKey:GraphID;references:ID;constraint:OnDelete:CASCADE;"`
	Cards       []DBEventCard  `gorm:"foreignKey:GraphID;references:ID;constraint:OnDelete:CASCADE;"`
	Connections []DBConnection `gorm:"foreignKey:GraphID;references:ID;constraint:OnDelete:CASCADE;"`
}

type DataGraph struct {
	Counter uint `json:"counter"`
}
