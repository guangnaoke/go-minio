package router

import (
	"minio_server/middleware"
	"minio_server/repositories"
	"minio_server/services"

	_ "minio_server/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	// 注册服务
	userRepository := repositories.NewUserManagerRepository("user", nil)
	userService := services.NewUserService(userRepository)

	bucketsService := services.NewBucketsService()
	objectService := services.NewObjectService()

	// 全局加入cors验证
	r.Use(middleware.Cors())

	user := r.Group("/api/user")
	{
		user.POST("/login", userService.Login)
		user.GET("/info", middleware.Auth(), userService.UserInfo)
	}

	bukets := r.Group("/api/buckets")
	bukets.Use(middleware.Auth())
	{
		bukets.GET("/list", bucketsService.List)
		bukets.GET("/exists", middleware.AuthCheckRole(), bucketsService.Exists)
		bukets.POST("/remove", middleware.AuthCheckRole(), bucketsService.Remove)
		bukets.GET("/listobjects", bucketsService.ListObjects)
	}

	object := r.Group("/api/object")
	object.Use(middleware.Auth())
	{
		object.GET("/stat", objectService.Stat)
		object.POST("/remove", middleware.AuthCheckRole(), objectService.Remove)
		object.POST("/upload", middleware.AuthCheckRole(), objectService.Put)
		object.GET("/url", middleware.AuthCheckRole(), objectService.GetObjectUrl)
	}

	// swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
