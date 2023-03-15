package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/model/vo"
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

func GetProvincePage(c *gin.Context) {
	provinceService := service.ProvinceService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := provinceService.GetProvincePage(c.Request.Context(), p)
		c.JSON(200, res)
	}
}

func GetProvincePageFuzzy(c *gin.Context) {
	provinceService := service.ProvinceService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := provinceService.GetProvincePageFuzzy(c.Request.Context(), p)
		c.JSON(200, res)
	}
}
