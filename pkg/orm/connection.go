package orm

type Connection struct {
	//ID
	ProjectId uint64
	Id        uint64 `gorm:"primarykey"`
	GraphId   uint64

	//Placement
	SourcePort string
	SourceId   uint64

	TargetPort string
	TargetId   uint64

	//Invokation
	Data string

	// UI
	Name string
	Ui   string
}
