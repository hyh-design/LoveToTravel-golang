package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ltt-gc/service"
)

func GetUserByEmail(c *gin.Context) {
	userService := service.UserService{}
	res := userService.GetUserByEmail(c.Request.Context(), c.Param("email"))
	c.JSON(200, res)
}

func GetUserList(c *gin.Context) {
	userService := service.UserService{}
	res := userService.GetUserList(c.Request.Context())
	c.JSON(200, res)
}

func CreateUser(c *gin.Context) {
	userService := service.UserService{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.CreateUser(c.Request.Context())
		c.JSON(200, res)
	}
}

func Login(c *gin.Context) {
	userService := service.UserService{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.Login(c.Request.Context())
		c.JSON(200, res)
	}
}

func GetUserByToken(c *gin.Context) {
	userService := service.UserService{}
	token := c.Request.Header.Get("Authorization")
	res := userService.GetUserByToken(c.Request.Context(), token)
	c.JSON(200, res)
}

func UpdateUser(c *gin.Context) {
	userService := service.UserService{}
	if err := c.ShouldBind(&userService); err == nil {
		res := userService.UpdateUser(c.Request.Context())
		fmt.Println("do put2")
		c.JSON(200, res)
	} else {
		fmt.Println(err)
	}
}

func DeleteUserById(c *gin.Context) {
	userService := service.UserService{}
	res := userService.DeleteUserById(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}
