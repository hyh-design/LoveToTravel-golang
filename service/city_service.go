package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/serializer"
)

type CityService struct {
	CityId       string
	CityName     string
	CityEname    string
	Lng          string
	Lat          string
	Url          string
	Introduction string
}

func (service *CityService) GetCityById(ctx context.Context, id string) serializer.Response {
	cityDao := dao.NewCityDao(ctx)
	city, err := cityDao.GetCityById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(city)
}

func (service *CityService) GetCityByName(ctx context.Context, name string) serializer.Response {
	cityDao := dao.NewCityDao(ctx)
	city, err := cityDao.GetCityByName(name)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(city)
}

func (service *CityService) GetCityList(ctx context.Context) serializer.Response {
	cityDao := dao.NewCityDao(ctx)
	cities, err := cityDao.GetCityList()
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(cities)
}
