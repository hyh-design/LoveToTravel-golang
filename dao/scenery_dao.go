package dao

import (
	"context"
	"gorm.io/gorm"
	"ltt-gc/config"
	"ltt-gc/model"
	"ltt-gc/model/vo"
)

type SceneryDao struct {
	*gorm.DB
}

func NewSceneryDao(ctx context.Context) *SceneryDao {
	return &SceneryDao{config.NewDBClient(ctx)}
}

func (dao *SceneryDao) CreateScenery(scenery *model.Scenery) (err error) {
	err = dao.DB.Create(&scenery).Debug().Error
	return
}

func (dao *SceneryDao) GetSceneryByName(name string) (scenery *model.Scenery, err error) {
	err = dao.DB.
		Where("name=?", name).First(&scenery).Debug().Error
	return
}

func (dao *SceneryDao) GetSceneryById(id string) (scenery *model.Scenery, err error) {
	err = dao.DB.
		Where("id=?", id).First(&scenery).Debug().Error
	return
}

func (dao *SceneryDao) GetSceneryPage(page vo.Page) (sceneries []*model.Scenery, err error) {
	err = dao.DB.Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Find(&sceneries).Error
	return
}

func (dao *SceneryDao) GetSceneryPageFuzzy(page vo.Page) (sceneries []*model.Scenery, err error) {
	err = dao.DB.Where("name like ?", "%"+page.QueryStr+"%").
		Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize).Debug().Find(&sceneries).Error
	return
}

func (dao *SceneryDao) Count() (total int, err error) {
	var t int64
	err = dao.DB.Model(&model.Scenery{}).Count(&t).Error
	total = int(t)
	return
}

func (dao *SceneryDao) GetSceneryList() (scenery []*model.Scenery, err error) {
	err = dao.DB.Find(&scenery).Debug().Error
	return
}
