package orm

type Graph struct {
	//ID
	ProjectId uint64
	Id        uint64 `gorm:"primarykey"`

	//Params
	LocalIdCounter uint64

	//UI
	Ui   string
	Name string

	//Relations
	Nodes       []*Node       `gorm:"constraint:OnDelete:CASCADE;"`
	Cards       []*EventCard  `gorm:"constraint:OnDelete:CASCADE;"`
	Connections []*Connection `gorm:"constraint:OnDelete:CASCADE;"`
}
