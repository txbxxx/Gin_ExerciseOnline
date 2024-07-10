/**
 * @Author tanchang
 * @Description //检查用户是否为普通权限
 * @Date 2024/7/9 16:33
 * @File:  isuser
 * @Software: GoLand
 **/

package middleware

import (
	"GinProject_ExerciseOnline/utils"
	"github.com/gin-gonic/gin"
)

func IsUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("token")
		claims, err := utils.AnalyseToken(auth)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "token认证失败",
			})
			c.Abort()
			return
		}
		if claims == nil {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		c.Set("user_claims", claims)
		c.Next()
	}
}
