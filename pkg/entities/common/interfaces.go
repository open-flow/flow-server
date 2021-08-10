package common

type ProjectObject interface {
	ByProject
	GetId() uint
}

type ByProject interface {
	GetProjectId() uint
}
