package model

import "gorm.io/gorm"

/**
 * gorm.Model对应数据库设计
 * ID - id自增长或赋值
 * CreatedAt、UpdatedAt、DeletedAt - created_at、update_at、deleted_at，datetime 3
 */

type Admin struct {
	gorm.Model
	ID       string
	Name     string
	Email    string
	Password string
}

func (v Admin) TableName() string {
	return "admin"
}
