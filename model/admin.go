package model

import "gorm.io/gorm"

/**
 * gorm.Model对应数据库设计
 * ID - id自增长或赋值
 * CreatedAt、UpdatedAt、DeletedAt - created_at、update_at、deleted_at，datetime 3
 */

type Admin struct {
	gorm.Model
	ID       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (v Admin) TableName() string {
	return "admin"
}
