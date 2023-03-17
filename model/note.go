package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	Id         primitive.ObjectID `form:"_id" json:"_id"`
	UserId     string             `form:"userId" json:"userId"`
	UserName   string             `form:"userName" json:"userName"`
	Title      string             `form:"title" json:"title"`
	PlanId     string             `form:"planId" json:"planId"`
	Url        string             `form:"url" json:"url"`
	Content    string             `form:"content" json:"content"`
	Comment    interface{}        `form:"comment" json:"comment"`
	View       interface{}        `form:"view" json:"view"`
	Star       interface{}        `form:"star" json:"star"`
	Trip       interface{}        `form:"trip" json:"trip"`
	Deleted    string             `form:"deleted" json:"deleted"`
	CreateTime string             `form:"createTime" json:"createTime"`
	UpdateTime string             `form:"updateTime" json:"updateTime"`
}
