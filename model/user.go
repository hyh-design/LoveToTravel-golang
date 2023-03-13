package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         string `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Url        string `form:"url" json:"url" gorm:"default:https://eemall.oss-cn-hangzhou.aliyuncs.com/Snipaste_2022-12-19_13-35-58.png"`
	Grade      string `form:"grade" json:"grade"`
	Experience string `form:"experience" json:"experience"`
	Tele       string `form:"tele" json:"tele"`
	Birthday   string `form:"birthday" json:"birthday"`
	Post       string `form:"post" json:"post"`
	Profession string `form:"profession" json:"profession"`
	Signature  string `form:"signature" json:"signature"`
	Gender     string `form:"gender" json:"gender"`
	Address    string `form:"address" json:"address"`
	Visits     int64  `form:"visits" json:"visits" gorm:"default:0"`
	Status     string `form:"status" json:"status" gorm:"default:0"`
}

func (v User) TableName() string {
	return "user"
}
