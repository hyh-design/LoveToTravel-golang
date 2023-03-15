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
		admin.POST("/login", api.LoginAdmin)

		admin.GET("/:email", api.GetAdminByEmail)
		admin.GET("/list", api.GetAdminList)
		admin.POST("", api.CreateAdmin)
		admin.PUT("", api.UpdateAdmin)
		admin.DELETE("/:id", api.DeleteAdminById)
	}

	/**
	TODO
	更改密码
	更改个人信息
	分页查询
	*/

	r.POST("/login", api.Login)
	r.POST("/register", api.CreateUser)

	user := r.Group("/user", handler.JWTAuthMiddleware())
	{
		user.GET("/token", api.GetUserByToken)
		user.GET("/:email", api.GetUserByEmail)
		user.GET("/list", api.GetUserList)
		user.PUT("", api.UpdateUser)
		user.DELETE("/:id", api.DeleteUserById)
	}

	city := r.Group("/city")
	{
		city.GET("", api.GetCityList)
		city.GET("/:id", api.GetCityById)
		city.GET("name/:name", api.GetCityByName)
	}

	province := r.Group("/province")
	{
		province.GET("", api.GetProvinceList)
		province.GET("/:id", api.GetProvinceById)
		province.GET("name/:name", api.GetProvinceByName)
	}

	return r
}
