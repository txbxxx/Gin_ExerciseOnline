/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 22:17
 * @File:  submitSvc
 * @Software: GoLand
 **/

package service

import (
	"GinProject_ExerciseOnline/dao"
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// SearchSubmitList 提交列表
// @Tags 公共方法
// @Summary 查询提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param user_identity query string false "user_identity"
// @Param problem_identity query string false "problem_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /submit/searchSubmitList [get]
func SearchSubmitList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户唯一标识，问题唯一标识、提交状态
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	userIdentity := c.Query("user_identity")
	problemIdentity := c.Query("problem_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//获取数据
	var count int64
	tx := dao.GetSubmitList(userIdentity, problemIdentity, status)
	submitList := make([]*model.Submit, 1)
	err := tx.Count(&count).Offset(page).Limit(size).Find(&submitList).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"page":        page,
			"size":        size,
			"count":       count,
			"submit_list": submitList,
		},
	})
}
