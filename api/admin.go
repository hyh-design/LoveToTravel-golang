package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ltt-gc/service"
)

//api——转化为json

func GetAdminByEmail(c *gin.Context) {
	adminService := service.AdminService{}
	res := adminService.GetAdminByEmail(c.Request.Context(), c.Param("email"))
	c.JSON(200, res)
}

func GetAdminList(c *gin.Context) {
	adminService := service.AdminService{}
	res := adminService.GetAdminList(c.Request.Context())
	c.JSON(200, res)
}

func CreateAdmin(c *gin.Context) {
	adminService := service.AdminService{}
	if err := c.ShouldBind(&adminService); err == nil {
		res := adminService.CreateAdmin(c.Request.Context())
		c.JSON(200, res)
	}
}

func LoginAdmin(c *gin.Context) {
	adminService := service.AdminService{}
	if err := c.ShouldBind(&adminService); err == nil {
		res := adminService.Login(c.Request.Context())
		c.JSON(200, res)
	}
}

func UpdateAdmin(c *gin.Context) {
	adminService := service.AdminService{}
	if err := c.ShouldBind(&adminService); err == nil {
		res := adminService.UpdateAdmin(c.Request.Context())
		fmt.Println("do put2")
		c.JSON(200, res)
	} else {
		fmt.Println(err)
	}
}

func DeleteAdminById(c *gin.Context) {
	adminService := service.AdminService{}
	res := adminService.DeleteAdminById(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}
