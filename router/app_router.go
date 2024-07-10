/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 1:04
 * @File:  app_router
 * @Software: GoLand
 **/

package router

import (
	_ "GinProject_ExerciseOnline/docs" //这里要引入项目内的docs文件
	"GinProject_ExerciseOnline/middleware"
	"GinProject_ExerciseOnline/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	httpServer := gin.Default()
	httpServer.Use(middleware.Cors())

	//路由规则

	httpServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // swagger 配置
	httpServer.GET("/ping", service.Ping)

	//问题接口
	problemApi := httpServer.Group("/problem")
	{
		problemApi.GET("/getProblemList", service.GetProblemList)
		problemApi.GET("/problemDetail", service.ProblemDetail)
	}

	//用户接口
	userApi := httpServer.Group("/user")
	{
		userApi.GET("/userDetail", service.UserDetail)
		userApi.POST("/login", service.Login)
		userApi.POST("/sendCode", service.SendCode)
		userApi.POST("/register", service.Register)
		userApi.GET("/ranking", service.UserRanking)
		userApi.POST("/submit", middleware.IsUser(), service.Submit)
	}

	//提交接口
	submitApi := httpServer.Group("/submit")
	{
		submitApi.GET("/searchSubmitList", service.SearchSubmitList)
	}

	//管理权限接口
	adminApi := httpServer.Group("/admin", middleware.IsAdmin())
	{
		adminApi.POST("/createProblem", service.CreateProblem)
		adminApi.GET("/getCategoryList", service.GetCategoryList)
		adminApi.DELETE("/delCategory", service.DelCategory)
		adminApi.POST("/createCategory", service.CreateCategory)
		adminApi.PUT("/modifyCategory", service.ModifyCategory)
	}

	return httpServer

}
