package service

import (
	"context"
	logging "github.com/sirupsen/logrus"
	"ltt-gc/dao"
	"ltt-gc/model"
	"ltt-gc/serializer"
	"ltt-gc/utils"
	"strconv"
)

type AdminService struct {
	ID       string `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// GetAdminByEmail
// @Tags admin-service
// @Router /admin/:email [get]
func (service *AdminService) GetAdminByEmail(ctx context.Context, email string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	admin, err := adminDao.GetAdminByEmail(email)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(admin)
}

// GetAdminList
// @Tags admin-service
// @Router /admin/list [get]
func (service *AdminService) GetAdminList(ctx context.Context) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	admins, err := adminDao.GetAdminList()
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(admins)
}

// CreateAdmin
// @Tags admin-service
// @Router /admin [post]
func (service *AdminService) CreateAdmin(ctx context.Context) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	snowFlake := utils.SnowFlake{}
	id := snowFlake.Generate()
	admin := &model.Admin{
		ID:       strconv.FormatInt(id, 10),
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err := adminDao.CreateAdmin(admin)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(admin)
}

func (service *AdminService) Login(ctx context.Context) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	admin := &model.Admin{
		Email:    service.Email,
		Password: service.Password,
	}
	err := adminDao.Login(admin)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(admin)
}

// UpdateAdmin
// @Tags admin-service
// @Router /admin [put]
func (service *AdminService) UpdateAdmin(ctx context.Context) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	admin := &model.Admin{
		ID:       service.ID,
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err := adminDao.UpdateAdmin(service.ID, admin)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(admin)
}

// DeleteAdminById
// @Tags admin-service
// @Router /admin/:id [delete]
func (service *AdminService) DeleteAdminById(ctx context.Context, id string) serializer.Response {
	adminDao := dao.NewAdminDao(ctx)
	err := adminDao.DeleteAdminById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(nil)
}
