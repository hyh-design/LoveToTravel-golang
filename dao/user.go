package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{config.NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Debug().Error
	return
}

func (dao *UserDao) Login(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("email=? and password = ?", user.Email, user.Password).
		First(&user).Debug().Error
	return
}

func (dao *UserDao) GetUserByEmail(email string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).
		Where("email=?", email).
		First(&user).Debug().Error
	return
}

func (dao *UserDao) GetUserList() (user []*model.User, err error) {
	err = dao.DB.Model(&model.User{}).Find(&user).Debug().Error
	return
}

func (dao *UserDao) UpdateUser(id string, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).
		Where("id=?", id).Updates(user).Debug().Error
	return
}

func (dao *UserDao) DeleteUserById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.User{}).Debug().Error
	return
}
