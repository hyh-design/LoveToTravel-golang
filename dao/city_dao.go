package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
	"ltt-gc/model/vo"
)

type CityDao struct {
	*gorm.DB
}

func NewCityDao(ctx context.Context) *CityDao {
	return &CityDao{config.NewDBClient(ctx)}
}

func (dao *CityDao) CreateCity(city *model.City) (err error) {
	err = dao.DB.Create(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityByName(name string) (city *model.City, err error) {
	err = dao.DB.Where("city_name=?", name).First(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityById(id string) (city *model.City, err error) {
	err = dao.DB.Where("city_id=?", id).First(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityList() (city []*model.City, err error) {
	err = dao.DB.Find(&city).Debug().Error
	return
}

func (dao *CityDao) GetCityPage(page vo.Page) (cities []*model.City, err error) {
	err = dao.DB.Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&cities).Error
	return
}

func (dao *CityDao) GetCityPageFuzzy(page vo.Page) (cities []*model.City, err error) {
	err = dao.DB.Where("city_name like ?", "%"+page.QueryStr+"%").
		Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Debug().Find(&cities).Error
	return
}

func (dao *CityDao) Count() (total int, err error) {
	var t int64
	err = dao.DB.Model(&model.City{}).Count(&t).Error
	total = int(t)
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
