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
	"GinProject_ExerciseOnline/utils"
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

// Login 用户登录
// @Tags 公共方法
// @Summary 用户登录
// @Param user_name formData  string false "user_name"
// @Param password formData  string false "password"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	//判断用户名或者密码是否为空
	if userName == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "用户名和密码为空",
		})
		return
	}

	//生成password的md5值
	password = utils.GetMd5(password)

	//获取用户信息
	var data model.User
	errData := define.DB.Model(&model.User{}).Where("name = ? and password = ?", userName, password).Find(&data).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(500, gin.H{
			"code": 500,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errData.Error(),
		})
		return
	}

	//保障再次判断用户是否存在
	if data == (model.User{}) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "用户名或密码错误",
		})
		return
	}

	//生成token
	token, errToken := utils.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "生成Token失败" + errToken.Error(),
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功！！",
		"data": gin.H{
			"token":    token,
			"is_admin": data.IsAdmin,
		},
	})
}
