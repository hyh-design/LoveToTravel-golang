package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/model/vo"
	"ltt-gc/service"
)

func GetSceneryById(c *gin.Context) {
	sceneryService := service.SceneryService{}
	res := sceneryService.GetSceneryById(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

func GetSceneryByName(c *gin.Context) {
	sceneryService := service.SceneryService{}
	res := sceneryService.GetSceneryByName(c.Request.Context(), c.Param("name"))
	c.JSON(200, res)
}

func GetSceneryList(c *gin.Context) {
	sceneryService := service.SceneryService{}
	res := sceneryService.GetSceneryList(c.Request.Context())
	c.JSON(200, res)
}

func GetSceneryPage(c *gin.Context) {
	sceneryService := service.SceneryService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := sceneryService.GetSceneryPage(c.Request.Context(), p)
		c.JSON(200, res)
	}
}

func GetSceneryPageFuzzy(c *gin.Context) {
	sceneryService := service.SceneryService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := sceneryService.GetSceneryPageFuzzy(c.Request.Context(), p)
		c.JSON(200, res)
	}
}
