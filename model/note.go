package model

type Note struct {
	UserId     string      `form:"userId" bson:"userId"`
	UserName   string      `form:"userName" bson:"userName"`
	Title      string      `form:"title" bson:"title"`
	PlanId     string      `form:"planId" bson:"planId"`
	Url        string      `form:"url" bson:"url"`
	Content    string      `form:"content" bson:"content"`
	Comment    interface{} `form:"comment" bson:"comment"`
	View       interface{} `form:"view" bson:"view"`
	Star       interface{} `form:"star" bson:"star"`
	Trip       interface{} `form:"trip" bson:"trip"`
	Deleted    string      `form:"deleted" bson:"deleted"`
	CreateTime string      `form:"createTime" bson:"createTime"`
	UpdateTime string      `form:"updateTime" bson:"updateTime"`
}
