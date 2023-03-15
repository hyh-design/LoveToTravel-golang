package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
)

type CityDao struct {
	*gorm.DB
}

func NewCityDao(ctx context.Context) *CityDao {
	return &CityDao{config.NewDBClient(ctx)}
}

func (dao *CityDao) CreateCity(city *model.City) (err error) {
	err = dao.DB.Model(&model.City{}).Create(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityByName(name string) (city *model.City, err error) {
	err = dao.DB.Model(&model.City{}).
		Where("city_name=?", name).
		First(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityById(id string) (city *model.City, err error) {
	err = dao.DB.Model(&model.City{}).
		Where("city_id=?", id).
		First(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityList() (city []*model.City, err error) {
	err = dao.DB.Model(&model.City{}).Find(&city).Debug().Error
	return
}

func (dao *CityDao) UpdateCity(id string, city *model.City) (err error) {
	err = dao.DB.Model(&model.City{}).
		Where("id=?", id).Updates(city).Debug().Error
	return
}

func (dao *CityDao) DeleteCityById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.City{}).Debug().Error
	return
}
