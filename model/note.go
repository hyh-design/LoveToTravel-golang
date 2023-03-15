package model

type Note struct {
	Id       string
	UserId   string
	UserName string
	Title    string
	PlanId   string
	Url      string
	Content  string
	Comment  int64
	View     int64
	Star     int64
	Trip     []string
}
