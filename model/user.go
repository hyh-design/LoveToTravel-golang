package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         string `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Url        string `form:"url" json:"url"`
	Grade      string `form:"grade" json:"grade"`
	Experience string `form:"experience" json:"experience"`
	Tele       string `form:"tele" json:"tele"`
	Birthday   string `form:"birthday" json:"birthday"`
	Post       string `form:"post" json:"post"`
	Profession string `form:"profession" json:"profession"`
	Signature  string `form:"signature" json:"signature"`
	Gender     string `form:"gender" json:"gender"`
	Address    string `form:"address" json:"address"`
	Visits     string `form:"visits" json:"visits"`
	Status     string `form:"status" json:"status"`
}

func (v User) TableName() string {
	return "user"
}
