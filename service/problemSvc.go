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
	"GinProject_ExerciseOnline/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

// GetProblemList 题目列表
// @Tags 公共方法
// @Summary 查询问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param category_identity query string false "category_identity"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem/getProblemList [get]
func GetProblemList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户输入搜素字段、分类唯一标识
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	searchData := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//总共数据条数接收
	var count int64

	//模糊搜索
	searchData = "%" + searchData + "%"

	//接收问题切片
	list := make([]*model.Problem, 0)
	//查询问题列表，并且限制分页
	tx := dao.GetProblemList(searchData, categoryIdentity)

	err := tx.Count(&count).Omit("context").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}
	//返回数据
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"count": count,
			"page":  page,
			"size":  size,
			"list":  list,
		},
	})
}

// ProblemDetail 题目详情
// @Tags 公共方法
// @Summary 查询问题详情
// @Param identity query string false "identity"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem/problemDetail [get]
func ProblemDetail(c *gin.Context) {
	//获取问题唯一标识
	problemIdentity := c.Query("identity")
	if problemIdentity == "" {
		c.JSON(200, gin.H{
			"code": -1,
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
				"code": -1,
				"msg":  "题目不存在",
			})
			return
		}
		c.JSON(500, gin.H{
			"code": -1,
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

// CreateProblem 创建题目
// @Tags 管理员方法
// @Summary 创建题目
// @Param token header string true "token"
// @Param title formData string true "title"
// @Param context formData string true "context"
// @Param max_mem formData int false "max_mem"
// @Param max_runtime formData int false "max_runtime"
// @Param category_name formData array  false "category_name"
// @Param input_case formData string true "input_case"
// @Param output_case formData string true "output_case"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /admin/createProblem [post]
func CreateProblem(c *gin.Context) {
	title := c.PostForm("title")
	context := c.PostForm("context")
	categoryName := c.PostFormArray("category_name")
	inputCase := c.PostForm("input_case")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMem, _ := strconv.Atoi(c.PostForm("max_mem"))
	outputCase := c.PostForm("output_case")
	if title == "" || context == "" || inputCase == "" || outputCase == "" {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}

	//查看问题分类是否存在，如果不存在则创建
	categoryData := make([]model.Category, 0)
	for _, v := range categoryName {
		var category model.Category
		err := define.DB.First(&category, "name = ?", v).Error
		if err != nil {
			//ErrRecordNotFound表示分类不存在
			if errors.Is(err, gorm.ErrRecordNotFound) {
				category = model.Category{
					Name:     v,
					Identity: utils.GenerateUUID(),
				}
				err := define.DB.Create(&category).Error
				if err != nil {
					c.JSON(200, gin.H{
						"code": -1,
						"msg":  "创建失败" + err.Error(),
					})
					return
				}
				categoryData = append(categoryData, category)
			} else {
				c.JSON(200, gin.H{
					"code": -1,
					"msg":  "分类处理失败" + err.Error(),
				})
				return
			}
		}
		categoryData = append(categoryData, category)
	}

	//查找测试案例是否存在
	var testCaseData []model.TestCase
	err2 := define.DB.Model(&model.TestCase{}).Where("input = ? and output = ?", inputCase, outputCase).Error
	if err2 != nil {
		if errors.Is(err2, gorm.ErrRecordNotFound) {
			testCase := model.TestCase{
				Identity: utils.GenerateUUID(),
				Input:    inputCase,
				Output:   outputCase,
			}
			testCaseData = append(testCaseData, testCase)
		} else {
			c.JSON(200, gin.H{
				"code": -1,
				"msg":  "案列处理失败" + err2.Error(),
			})
			return
		}
	}

	//创建问题对象
	problem := &model.Problem{
		Identity:   utils.GenerateUUID(),
		Title:      title,
		Context:    context,
		MaxMem:     maxMem,
		MaxRuntime: maxRuntime,
		Category:   categoryData,
		TestCase:   testCaseData,
	}

	//创建问题
	err := define.DB.Create(&problem).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "创建失败" + err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})

}
