package module

type DBModule struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	DataHttpModule `gorm:"embed"`
}

type DeleteRequest struct {
	ID uint `json:"id"`
}

type DeleteResponse struct {
}

type ListRequest struct {
}

type ListResponse struct {
	Modules []DBModule
}

type DataHttpModule struct {
	Uri    string `json:"uri"`
	Module string `gorm:"uniq" json:"module"`
}

type ModuleDef struct {
	Module    string         `json:"module"`
	Functions []*FunctionDef `json:"functions"`
}

type FunctionDef struct {
	Function string `json:"function"`
}
