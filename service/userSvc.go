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
	"strconv"
	"time"
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
			"code": -1,
			"msg":  "用户名或密码为空",
		})
		return
	}

	//生成password的md5值
	password = utils.GetMd5(password)

	//获取用户信息
	var data model.User
	var cunt int64
	errData := define.DB.Model(&model.User{}).Where("name = ? and password = ?", userName, password).Count(&cunt).Find(&data).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误",
			})
			return
		}
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errData.Error(),
		})
		return
	}

	//保障再次判断用户是否存在
	if cunt > 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户名或密码错误",
		})
		return
	}

	//生成token
	token, errToken := utils.GenerateToken(data.Identity, data.Name, data.IsAdmin)
	if errToken != nil {
		c.JSON(200, gin.H{
			"code": "-1",
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

// SendCode 向邮箱发送验证码
// 可优化的点: 随机验证码
// @Tags 公共方法
// @Summary 向邮箱发送验证码
// @Param email formData  string false "email"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/sendCode [post]
func SendCode(c *gin.Context) {
	//获取发送的邮箱
	email := c.PostForm("email")
	if email == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "邮箱为空",
		})
		return
	}

	//获取随机验证码
	code := utils.GenerateCode()

	//将验证码存入redis
	define.RDB.Set(c, email, code, time.Second*300)

	//发送验证码
	err := utils.SendEmail(email, code)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "发送验证码失败" + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发送验证码成功",
	})
}

// Register 注册用户
// 可优化的点: 随机验证码
// @Tags 公共方法
// @Summary 注册用户
// @Param user_name formData  string false "user_name"
// @Param password formData  string false "password"
// @Param email formData  string false "email"
// @Param phone formData  string false "phone"
// @Param code formData  string false "code"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	//获取用户填写的信息
	userName := c.PostForm("user_name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	userCode := c.PostForm("code")

	if userName == "" || password == "" || email == "" || phone == "" || userCode == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数错误，请检查",
		})
		return
	}

	//查找是否存在用户
	var userData model.User
	errSearchUser := define.DB.Model(&model.User{}).Where("name = ? ", userName).Find(&userData).Error
	if errSearchUser != nil {
		c.JSON(500, gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errSearchUser.Error(),
		})
		return
	}
	if userData != (model.User{}) {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}

	//接收验证码
	sysCode, err := define.RDB.Get(c, email).Result()
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "验证码错误,请重新获取验证码",
		})
		return
	}
	//对比验证码
	if sysCode != userCode {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "验证码不匹配！！",
		})
	}

	//判断邮箱是否存在
	errSearchUser = define.DB.Model(&model.User{}).Where("mail = ? ", email).Find(&userData).Error
	if errSearchUser != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errSearchUser.Error(),
		})
		return
	}

	if userData != (model.User{}) {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "邮箱已经被注册",
		})
		return
	}

	//设置用户唯一标识
	identity := utils.GenerateUUID()

	//生成密码的md值
	password = utils.GetMd5(password)

	//创建用户对象
	user := &model.User{
		Name:     userName,
		Password: password,
		Identity: identity,
		Phone:    phone,
		Mail:     email,
	}
	//插入数据库
	errInsert := define.DB.Model(&model.User{}).Create(&user).Error
	if errInsert != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + errInsert.Error(),
		})
		return
	}

	//生成token
	token, errToken := utils.GenerateToken(user.Identity, user.Name, user.IsAdmin)
	if errToken != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "生成Token失败" + errToken.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功！！",
		"data": gin.H{
			"token": token,
		},
	})
}

// UserRanking 用户排名
// 可优化的点: 随机验证码
// @Tags 公共方法
// @Summary 用户排名
// @Param page formData  string false "page"
// @Param size formData  string false "size"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/ranking [get]
func UserRanking(c *gin.Context) {
	//获取页面也页面数据个数
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//搜素
	var count int64

	users := make([]*model.User, 0)
	//排序查找用户
	err := define.DB.Model(&model.User{}).Count(&count).Order("finish_problem_num desc,submit_problem_num asc").Offset(page).Limit(size).Find(&users).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "服务器内部错误，请联系管理员或者提交issue" + err.Error(),
		})
		return
	}

	//成功返回
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": gin.H{
			"count": count,
			"list":  users,
		},
	})
}
