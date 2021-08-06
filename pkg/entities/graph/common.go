package graph

type DataUI struct {
	Name string
	Ui   string
}

type IDGraph struct {
	IDProject
	GraphID uint `gorm:"index"`
}

type IDProject struct {
	ProjectID uint
	ID        uint `gorm:"primaryKey"`
}
