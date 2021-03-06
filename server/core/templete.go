package core

// Templete ...
type Templete struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required,max=40"`
	Content  string `json:"content" validate:"required"`
	Language string `json:"language" validate:"max=10"`
	DataBase string `json:"data_base" validate:"max=10"`
	Orm      string `json:"orm" validate:"max=10"`
	Type     string `json:"type" validate:"required,max=10"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

// TempleteQuery 分页查询参数
type TempleteQuery struct {
	Pager
	Name     string `json:"name"`
	Language string `json:"language"`
	DataBase string `json:"data_base"`
	Orm      string `json:"orm"`
	Type     string `json:"type"`
	PID      int64  `json:"pid"`
}

// TempleteLoadReq ...
type TempleteLoadReq struct {
	TempleteID int64 `json:"templete_id"`
	TID        int64 `json:"tid"`
}

// TempleteStore ...
type TempleteStore interface {
	Create(*Templete) error

	Update(*Templete) error

	FindID(int64) (*Templete, error)

	FindName(string) (*Templete, error)

	List(*TempleteQuery) ([]*Templete, int, error)

	Delete(int64) error

	CreateProjectTempleteRelation(pid, tid int64) error

	DeleteProjectTempleteRelationByPID(pid int64) error
}
