package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
	"ltt-gc/model/vo"
)

type ProvinceDao struct {
	*gorm.DB
}

func NewProvinceDao(ctx context.Context) *ProvinceDao {
	return &ProvinceDao{config.NewDBClient(ctx)}
}

func (dao *ProvinceDao) CreateProvince(province *model.Province) (err error) {
	err = dao.DB.Create(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceByName(name string) (province *model.Province, err error) {
	err = dao.DB.Where("province_name=?", name).First(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceById(id string) (province *model.Province, err error) {
	err = dao.DB.Where("province_id=?", id).First(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvinceList() (province []*model.Province, err error) {
	err = dao.DB.Find(&province).Debug().Error
	return
}

func (dao *ProvinceDao) GetProvincePage(page vo.Page) (provinces []*model.Province, err error) {
	err = dao.DB.Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&provinces).Error
	return
}

func (dao *ProvinceDao) GetProvincePageFuzzy(page vo.Page) (provinces []*model.Province, err error) {
	err = dao.DB.Where("province_name like ?", "%"+page.QueryStr+"%").
		Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Debug().Find(&provinces).Error
	return
}

func (dao *ProvinceDao) Count() (total int, err error) {
	var t int64
	err = dao.DB.Model(&model.Province{}).Count(&t).Error
	total = int(t)
	return
}

func (dao *ProvinceDao) UpdateProvince(id string, province *model.Province) (err error) {
	err = dao.DB.Where("id=?", id).Updates(province).Debug().Error
	return
}

func (dao *ProvinceDao) DeleteProvinceById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&model.Province{}).Debug().Error
	return
}
