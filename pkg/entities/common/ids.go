package common

type IDProject struct {
	ProjectId uint `json:"projectId,omitempty" form:"projectId"`
	Id        uint `gorm:"primaryKey,omitempty" json:"id,omitempty" form:"id"`
}

var _ ProjectObject = (*IDProject)(nil)

func (p *IDProject) GetId() uint {
	return p.Id
}

func (p *IDProject) GetProjectId() uint {
	return p.ProjectId
}

type ByProjectId struct {
	ProjectId uint `json:"projectId" form:"projectId"`
}

var _ ByProject = (*ByProjectId)(nil)

func (b *ByProjectId) GetProjectId() uint {
	return b.ProjectId
}
