package graph

type DBGraph struct {
	IDProject

	DataUI
	DataGraph

	Nodes       []DBNode       `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"nodes,omitempty"`
	Cards       []DBEventCard  `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"cards,omitempty"`
	Connections []DBConnection `gorm:"foreignKey:GraphId;references:Id;constraint:OnDelete:CASCADE;" json:"connections,omitempty"`
}

type DataGraph struct {
	Counter uint `json:"counter"`
}
