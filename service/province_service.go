package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/serializer"
)

type ProvinceService struct {
	ProvinceId   string
	ProvinceName string
	Url          string
	Introduction string
}

func (service *ProvinceService) GetProvinceById(ctx context.Context, id string) serializer.Response {
	provinceDao := dao.NewProvinceDao(ctx)
	province, err := provinceDao.GetProvinceById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(province)
}

func (service *ProvinceService) GetProvinceByName(ctx context.Context, name string) serializer.Response {
	provinceDao := dao.NewProvinceDao(ctx)
	province, err := provinceDao.GetProvinceByName(name)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(province)
}

func (service *ProvinceService) GetProvinceList(ctx context.Context) serializer.Response {
	provinceDao := dao.NewProvinceDao(ctx)
	provinces, err := provinceDao.GetProvinceList()
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	return serializer.Success(provinces)
}
