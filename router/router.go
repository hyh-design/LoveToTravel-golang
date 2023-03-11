package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
	"ltt-gc/api"
	"ltt-gc/config"
	"ltt-gc/docs"
	"ltt-gc/handler"
)

var Db *gorm.DB

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(config.Cors())

	r.GET("/home", handler.JWTAuthMiddleware(), handler.HomeHandler)

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", func(context *gin.Context) {
		fmt.Println("success")
		context.JSON(200, "success")
	})

	admin := r.Group("/admin")
	{
		admin.GET("/:email", api.GetAdminByEmail)
		admin.GET("/list", api.GetAdminList)
		admin.POST("", api.CreateAdmin)
		admin.POST("/login", api.LoginAdmin)
		admin.PUT("", api.UpdateAdmin)
		admin.DELETE("/:id", api.DeleteAdminById)
	}

	user := r.Group("/user")
	{
		user.POST("/login", api.Login)
	}

	return r
}
