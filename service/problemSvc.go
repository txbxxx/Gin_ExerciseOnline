/**
 * @Author tanchang
 * @Description //问题的逻辑处理
 * @Date 2024/7/4 14:12
 * @File:  problemSvc
 * @Software: GoLand
 **/

package service

import (
	"GinProject_ExerciseOnline/dao"
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// SearchProblemList 题目列表
// @Tags 公共方法
// @Summary 查询问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param category_identity query string false "category_identity"
// @Param searchData query string false "searchData"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem/searchProblemList [get]
func SearchProblemList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户输入搜素字段、分类唯一标识
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	searchData := c.Query("searchData")
	categoryIdentity := c.Query("category_identity")

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//总共数据条数接收
	var count int64

	//模糊搜索
	searchData = "%" + searchData + "%"

	//接收问题切片
	problemList := make([]*model.Problem, 0)
	//查询问题列表，并且限制分页
	tx := dao.GetProblemList(searchData, categoryIdentity)

	err := tx.Count(&count).Omit("context").Offset(page).Limit(size).Find(&problemList).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "查询失败",
		})
		return
	}
	//返回数据
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"count":        count,
			"page":         page,
			"size":         size,
			"problem_list": problemList,
		},
	})
}

// ProblemDetail 题目详情
// @Tags 公共方法
// @Summary 查询问题详情
// @Param problem_identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem/problemDetail [get]
func ProblemDetail(c *gin.Context) {
	//获取问题唯一标识
	problemIdentity := c.Query("problem_identity")
	if problemIdentity == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "问题标识为空！",
		})
		return
	}

	var problemData model.Problem
	// 获取问题详细信息
	err := define.DB.Preload("Category").Find(&problemData, "identity = ?", problemIdentity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "题目不存在",
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
		"data": problemData,
	})
}
