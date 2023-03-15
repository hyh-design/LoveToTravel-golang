package api

import (
	"github.com/gin-gonic/gin"
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
