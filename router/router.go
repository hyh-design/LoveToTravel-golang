package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.POST("/login", api.Login)
	r.POST("/register", api.CreateUser)

	user := r.Group("/user", handler.JWTAuthMiddleware())
	//user := r.Group("/user")
	{
		user.GET("/token", api.GetUserByToken)
		user.GET("/:email", api.GetUserByEmail)
		user.GET("/list", api.GetUserList)
		user.PUT("", api.UpdateUser)
		user.DELETE("/:id", api.DeleteUserById)
		user.POST("/page", api.GetUserPage)
		user.POST("/page/query", api.GetUserPageFuzzy)
	}

	city := r.Group("/city")
	{
		city.GET("", api.GetCityList)
		city.GET("/:id", api.GetCityById)
		city.GET("/name/:name", api.GetCityByName)
		city.POST("/page", api.GetCityPage)
		city.POST("/page/query", api.GetCityPageFuzzy)
	}

	province := r.Group("/province")
	{
		province.GET("", api.GetProvinceList)
		province.GET("/:id", api.GetProvinceById)
		province.GET("/name/:name", api.GetProvinceByName)
		province.POST("/page", api.GetProvincePage)
		province.POST("/page/query", api.GetProvincePageFuzzy)
	}

	scenery := r.Group("/scenery")
	{
		scenery.GET("", api.GetSceneryList)
		scenery.GET("/:id", api.GetSceneryById)
		scenery.GET("/name/:name", api.GetSceneryByName)
		scenery.POST("/page", api.GetSceneryPage)
		scenery.POST("/page/query", api.GetSceneryPageFuzzy)
	}

	note := r.Group("/note")
	{
		note.GET("/:id", api.GetNoteById)
		note.GET("/", api.GetNoteList)
		note.POST("/page", api.GetNotePage)
		note.POST("/page/query", api.GetNotePageFuzzy)
		note.POST("/", api.CreateNote)
		note.PUT("/", api.UpdateNote)
		note.DELETE("/:id", api.DeleteNoteById)
	}

	plan := r.Group("/plan")
	{
		plan.GET("/:id", api.GetPlanById)
		plan.GET("/", api.GetPlanList)
		plan.POST("/page", api.GetPlanPage)
		plan.POST("/page/query", api.GetPlanPageFuzzy)
		plan.POST("/", api.CreatePlan)
		plan.PUT("/", api.UpdatePlan)
		plan.DELETE("/:id", api.DeletePlanById)
	}

	return r
}
