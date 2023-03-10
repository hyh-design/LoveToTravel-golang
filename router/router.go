package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
	"ltt-gc/api"
	"ltt-gc/config"
	"ltt-gc/docs"
	"ltt-gc/service"
)

var Db *gorm.DB

func NewRouter() *gin.Engine {

	r := gin.Default()
	r.Use(config.Cors())
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)

	r.GET("/init", func(context *gin.Context) {
		config.Init()
	})

	v1 := r.Group("/admin")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.GET("/:email", api.GetAdminByEmail)
		v1.GET("/list", api.GetAdminList)
		v1.POST("", api.CreateAdmin)
		v1.POST("/login", api.LoginAdmin)
		v1.PUT("", api.UpdateAdmin)
		v1.DELETE("/:id", api.DeleteAdminById)
	}

	return r
}
