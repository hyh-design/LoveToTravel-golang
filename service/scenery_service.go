package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/model/vo"
	"ltt-gc/serializer"
)

type SceneryService struct {
	Id           string
	Name         string
	Introduction string
	Score        string
	Ticket       string
	Opening      string
	Lng          string
	Lat          string
	Level        string
	Address      string
	Season       string
	Tips         string
	CityName     string
	CityId       string
	Url          string
	Tele         string
	Cluster      string
}

func (service *SceneryService) GetSceneryById(ctx context.Context, id string) serializer.Response {
	sceneryDao := dao.NewSceneryDao(ctx)
	scenery, err := sceneryDao.GetSceneryById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(scenery)
}

func (service *SceneryService) GetSceneryByName(ctx context.Context, name string) serializer.Response {
	sceneryDao := dao.NewSceneryDao(ctx)
	scenery, err := sceneryDao.GetSceneryByName(name)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(scenery)
}

func (service *SceneryService) GetSceneryList(ctx context.Context) serializer.Response {
	sceneryDao := dao.NewSceneryDao(ctx)
	sceneries, err := sceneryDao.GetSceneryList()
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(sceneries)
}

func (service *SceneryService) GetSceneryPage(ctx context.Context, p vo.Page) serializer.Response {
	sceneryDao := dao.NewSceneryDao(ctx)
	sceneries, err := sceneryDao.GetSceneryPage(p)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	p.Total, _ = sceneryDao.Count()
	p.PageSum = p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		p.PageNum++
	}
	p.Records = sceneries
	return serializer.Success(p)
}

func (service *SceneryService) GetSceneryPageFuzzy(ctx context.Context, p vo.Page) serializer.Response {
	sceneryDao := dao.NewSceneryDao(ctx)
	sceneries, err := sceneryDao.GetSceneryPageFuzzy(p)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	p.Total, _ = sceneryDao.Count()
	p.PageSum = p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		p.PageNum++
	}
	p.Records = sceneries
	return serializer.Success(p)
}
