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
	"GinProject_ExerciseOnline/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	httpServer := gin.Default()

	//路由规则

	httpServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // swagger 配置
	httpServer.GET("/ping", service.Ping)

	//问题接口
	problemApi := httpServer.Group("/problem")
	{
		problemApi.GET("/searchProblemList", service.SearchProblemList)
		problemApi.GET("/problemDetail", service.ProblemDetail)
	}

	//用户接口
	userApi := httpServer.Group("/user")
	{
		userApi.GET("/userDetail", service.UserDetail)
		userApi.POST("/login", service.Login)
		userApi.POST("/sendCode", service.SendCode)
		userApi.POST("/register", service.Register)
	}

	//提交接口
	submitApi := httpServer.Group("/submit")
	{
		submitApi.GET("/searchSubmitList", service.SearchSubmitList)
	}

	return httpServer

}
