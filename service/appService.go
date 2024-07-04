/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 4:11
 * @File:  appService
 * @Software: GoLand
 **/

package service

import (
	"github.com/gin-gonic/gin"
)

// Ping
// @Tags 公共方法
// @Summary 测试连接
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": "测试页面，看到此页面表示Gin启动成功！",
	})
}
