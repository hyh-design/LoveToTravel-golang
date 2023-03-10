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

	v1 := r.Group("/user")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})

		v1.GET("/:email", api.GetUserByEmail)
		v1.GET("/list", api.GetUserList)
		v1.POST("", api.CreateUser)
		v1.POST("/login", api.Login)
		v1.PUT("", api.UpdateUser)
		v1.DELETE("/:id", api.DeleteUserById)
	}

	return r
}
