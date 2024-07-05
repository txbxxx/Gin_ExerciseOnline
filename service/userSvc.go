/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 21:49
 * @File:  userSvc
 * @Software: GoLand
 **/

package service

import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserDetail 用户详细信息
// @Tags 公共方法
// @Summary 查用户详细信息
// @Param user_identity query string false "user_identity"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/userDetail [get]
func UserDetail(c *gin.Context) {
	//获取问题唯一标识
	userIdentity := c.Query("user_identity")
	if userIdentity == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "用户标识为空！",
		})
		return
	}

	var userData model.User
	// 获取问题详细信息
	err := define.DB.Omit("password").Find(&userData, "identity = ?", userIdentity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "用户不存在",
			})
			return
		}
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + err.Error(),
		})
		return
	}

	//获取成功返回查找到的数据
	c.JSON(200, gin.H{
		"code": 200,
		"data": userData,
	})
}
