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

type UserService struct {
	ID         string
	Name       string
	Email      string
	Password   string
	Url        string
	Grade      string
	Experience string
	Tele       string
	Birthday   string
	Post       string
	Profession string
	Signature  string
	Gender     string
	Address    string
	Visits     string
	Status     string
}

// GetUserByEmail
// @Tags user-service
// @Router /user/:email [get]
func (service *UserService) GetUserByEmail(ctx context.Context, email string) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByEmail(email)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// GetUserList
// @Tags user-service
// @Router /user/list [get]
func (service *UserService) GetUserList(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	users, err := userDao.GetUserList()
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(users)
}

// CreateUser
// @Tags user-service
// @Router /user [post]
func (service *UserService) CreateUser(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	isExist, _ := userDao.GetUserById(service.Email)
	if isExist.ID != "" {
		return serializer.Error()
	}
	snowFlake := utils.SnowFlake{}
	id := snowFlake.Generate()
	user := &model.User{
		ID:       strconv.FormatInt(id, 10),
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err := userDao.CreateUser(user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

func (service *UserService) Login(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	user := &model.User{
		Email:    service.Email,
		Password: service.Password,
	}
	err := userDao.Login(user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	token, _ := utils.GenToken(service.Email)
	return serializer.Success(token)
}
func (service *UserService) GetUserByToken(ctx context.Context, token string) serializer.Response {
	// token无效会经过拦截器处理, 仅解析token即可
	mc, _ := utils.ParseToken(token)
	email := mc.Email
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserByEmail(email)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// UpdateUser
// @Tags user-service
// @Router /user [put]
func (service *UserService) UpdateUser(ctx context.Context) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	_, err := userDao.GetUserById(service.ID)
	if err != nil {
		return serializer.Error()
	}
	user := &model.User{
		ID:       service.ID,
		Name:     service.Name,
		Email:    service.Email,
		Password: service.Password,
	}
	err = userDao.UpdateUser(service.ID, user)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(user)
}

// DeleteUserById
// @Tags user-service
// @Router /user/:id [delete]
func (service *UserService) DeleteUserById(ctx context.Context, id string) serializer.Response {
	userDao := dao.NewUserDao(ctx)
	err := userDao.DeleteUserById(id)
	if err != nil {
		logging.Info(err)
		return serializer.Error()
	}
	return serializer.Success(nil)
}
