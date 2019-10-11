package models

type Menu struct {
	ParentID int
	IsMenu   int // 1-模块 2-接口
	Name     string
	Remarks  string
	Icon     string
	URL      string
	Status   int
}
