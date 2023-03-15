package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/service"
)

func GetProvinceById(c *gin.Context) {
	provinceService := service.ProvinceService{}
	res := provinceService.GetProvinceById(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

func GetProvinceByName(c *gin.Context) {
	provinceService := service.ProvinceService{}
	res := provinceService.GetProvinceByName(c.Request.Context(), c.Param("name"))
	c.JSON(200, res)
}

func GetProvinceList(c *gin.Context) {
	provinceService := service.ProvinceService{}
	res := provinceService.GetProvinceList(c.Request.Context())
	c.JSON(200, res)
}
