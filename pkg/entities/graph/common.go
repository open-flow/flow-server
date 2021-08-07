package graph

type DataUI struct {
	Name string `json:"name,omitempty"`
	Ui   string `json:"ui,omitempty"`
}

type IDGraph struct {
	IDProject
	GraphId uint `gorm:"index" json:"graphId,omitempty"`
}

type IDProject struct {
	ProjectId uint `json:"projectId,omitempty" form:"projectId"`
	Id        uint `gorm:"primaryKey,omitempty" json:"id,omitempty" form:"id"`
}
