package graph

type DataUI struct {
	Name string `json:"name,omitempty"`
	Ui   string `json:"ui,omitempty"`
}

type IDGraph struct {
	IDProject
	GraphID uint `gorm:"index" json:"graphId,omitempty"`
}

type IDProject struct {
	ProjectID uint `json:"projectId,omitempty"`
	ID        uint `gorm:"primaryKey" json:"id,omitempty"`
}
