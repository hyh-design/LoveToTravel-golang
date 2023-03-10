package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
)

type AdminDao struct {
	*gorm.DB
}

func NewAdminDao(ctx context.Context) *AdminDao {
	return &AdminDao{config.NewDBClient(ctx)}
}

func (dao *AdminDao) CreateAdmin(admin *model.Admin) (err error) {
	err = dao.DB.Model(&model.Admin{}).Create(&admin).Debug().Error
	return
}

func (dao *AdminDao) Login(admin *model.Admin) (err error) {
	err = dao.DB.Model(&model.Admin{}).
		Where("email=? and password = ?", admin.Email, admin.Password).
		First(&admin).Debug().Error
	return
}

func (dao *AdminDao) GetAdminByEmail(email string) (admin *model.Admin, err error) {
	err = dao.DB.Model(&model.Admin{}).
		Where("email=?", email).
		First(&admin).Debug().Error
	return
}

func (dao *AdminDao) GetAdminList() (admin []*model.Admin, err error) {
	err = dao.DB.Model(&model.Admin{}).Find(&admin).Debug().Error
	return
}

func (dao *AdminDao) UpdateAdmin(id string, admin *model.Admin) (err error) {
	err = dao.DB.Model(&model.Admin{}).
		Where("id=?", id).Updates(admin).Debug().Error
	return
}

func (dao *AdminDao) DeleteAdminById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Admin{}).Debug().Error
	return
}
