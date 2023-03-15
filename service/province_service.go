package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/model/vo"
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

func (service *ProvinceService) GetProvincePage(ctx context.Context, p vo.Page) serializer.Response {
	provinceDao := dao.NewProvinceDao(ctx)
	provinces, err := provinceDao.GetProvincePage(p)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	p.Total, _ = provinceDao.Count()
	p.PageSum = p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		p.PageNum++
	}
	p.Records = provinces
	return serializer.Success(p)
}

func (service *ProvinceService) GetProvincePageFuzzy(ctx context.Context, p vo.Page) serializer.Response {
	provinceDao := dao.NewProvinceDao(ctx)
	provinces, err := provinceDao.GetProvincePageFuzzy(p)
	if err != nil {
		logging.Info(err)
		return serializer.Error(serializer.ServerError)
	}
	p.Total, _ = provinceDao.Count()
	p.PageSum = p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		p.PageNum++
	}
	p.Records = provinces
	return serializer.Success(p)
}
