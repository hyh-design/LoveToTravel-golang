package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
)

type ProvinceDao struct {
	*gorm.DB
}

func NewProvinceDao(ctx context.Context) *ProvinceDao {
	return &ProvinceDao{config.NewDBClient(ctx)}
}

func (dao *ProvinceDao) CreateProvince(province *model.Province) (err error) {
	err = dao.DB.Model(&model.Province{}).Create(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceByName(name string) (province *model.Province, err error) {
	err = dao.DB.Model(&model.Province{}).
		Where("province_name=?", name).
		First(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceById(id string) (province *model.Province, err error) {
	err = dao.DB.Model(&model.Province{}).
		Where("province_id=?", id).
		First(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceList() (province []*model.Province, err error) {
	err = dao.DB.Model(&model.Province{}).Find(&province).Debug().Error
	return
}

func (dao *ProvinceDao) UpdateProvince(id string, province *model.Province) (err error) {
	err = dao.DB.Model(&model.Province{}).
		Where("id=?", id).Updates(province).Debug().Error
	return
}

func (dao *ProvinceDao) DeleteProvinceById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Province{}).Debug().Error
	return
}
