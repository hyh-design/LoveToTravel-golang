package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/model/vo"
	"ltt-gc/service"
)

func GetPlanById(c *gin.Context) {
	planService := service.Plan{}
	res := planService.GetPlanById(c.Param("id"))
	c.JSON(200, res)
}

func GetPlanByCityId(c *gin.Context) {
	planService := service.Plan{}
	res := planService.GetPlanByCityId(c.Param("id"))
	c.JSON(200, res)
}

func GetPlanList(c *gin.Context) {
	planService := service.Plan{}
	res := planService.GetPlanList()
	c.JSON(200, res)
}

func GetPlanPage(c *gin.Context) {
	planService := service.Plan{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := planService.GetPlanPage(p)
		c.JSON(200, res)
	}
}

func GetPlanPageFuzzy(c *gin.Context) {
	planService := service.Plan{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := planService.GetPlanPageFuzzy(p)
		c.JSON(200, res)
	}
}

func CreatePlan(c *gin.Context) {
	planService := service.Plan{}
	if err := c.ShouldBind(&planService); err == nil {
		res := planService.CreatePlan()
		c.JSON(200, res)
	}
}

func UpdatePlan(c *gin.Context) {
	planService := service.Plan{}
	if err := c.ShouldBind(&planService); err == nil {
		res := planService.UpdatePlan()
		c.JSON(200, res)
	}
}

func DeletePlanById(c *gin.Context) {
	planService := service.Plan{}
	res := planService.DeletePlanById(c.Param("id"))
	c.JSON(200, res)
}
