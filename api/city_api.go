package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/model/vo"
	"ltt-gc/service"
)

func GetCityById(c *gin.Context) {
	cityService := service.CityService{}
	res := cityService.GetCityById(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

func GetCityByName(c *gin.Context) {
	cityService := service.CityService{}
	res := cityService.GetCityByName(c.Request.Context(), c.Param("name"))
	c.JSON(200, res)
}

func GetCityList(c *gin.Context) {
	cityService := service.CityService{}
	res := cityService.GetCityList(c.Request.Context())
	c.JSON(200, res)
}

func GetHotCity(c *gin.Context) {
	cityService := service.CityService{}
	res := cityService.GetHotCity(c.Request.Context())
	c.JSON(200, res)
}

func GetCityPage(c *gin.Context) {
	cityService := service.CityService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := cityService.GetCityPage(c.Request.Context(), p)
		c.JSON(200, res)
	}
}

func GetCityPageFuzzy(c *gin.Context) {
	cityService := service.CityService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := cityService.GetCityPageFuzzy(c.Request.Context(), p)
		c.JSON(200, res)
	}
}
