/**
 * @Author tanchang
 * @Description //问题的逻辑处理
 * @Date 2024/7/4 14:12
 * @File:  problemSvc
 * @Software: GoLand
 **/

package service

import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// SearchProblemList
// @Tags 公共方法
// @Summary 查询问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param searchData query string false "searchData"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem/searchProblemList [get]
func SearchProblemList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户输入搜素字段
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	searchData := c.Query("searchData")

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//总共数据条数接收
	var count int64

	//模糊搜索
	searchData = "%" + searchData + "%"

	//接收问题切片
	problemList := make([]*model.Problem, 0)
	//查询问题列表，并且限制分页
	err := define.DB.Model(&model.Problem{}).Count(&count).Offset(page).Limit(size).Where("title like ? OR context like ?", searchData, searchData).Find(&problemList).Error
	if err != nil {
		log.Println("搜索失败", err)
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
