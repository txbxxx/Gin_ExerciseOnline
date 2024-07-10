/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/8 16:16
 * @File:  categorySvc
 * @Software: GoLand
 **/

package service

import "C"
import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"GinProject_ExerciseOnline/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetCategoryList 分类列表
// @Tags 管理员方法
// @Summary 查询分类列表
// @Param token header string false "token"
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword  query string false "keyword"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /admin/getCategoryList [get]
func GetCategoryList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户输入搜素字段、分类唯一标识
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	searchData := c.Query("keyword")

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//总共数据条数接收
	var count int64

	//模糊搜索
	searchData = "%" + searchData + "%"

	//获取题目列表
	var categoryList []model.Category
	err := define.DB.Model(&model.Category{}).Where("name like ?", searchData).Count(&count).Offset(page).Limit(size).Find(&categoryList).Error
	if err != nil {
		log.Println("获取题目列表失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "获取题目列表失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"list":  categoryList,
			"count": count,
		},
	})
}

// CreateCategory 创建分类
// @Tags 管理员方法
// @Summary 创建分类
// @Param token header string true "token"
// @Param name formData string true "name"
// @Param parent_id formData string true "parent_id"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /admin/createCategory [post]
func CreateCategory(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))

	//创建唯一标识
	identity := utils.GenerateUUID()

	//查找分类是否存在
	var cnt int64
	err := define.DB.Model(&model.Category{}).Where("identity = ?", identity).Count(&cnt).Error
	if err != nil {
		log.Println("查找分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查找分类失败" + err.Error(),
		})
		return
	}

	if cnt > 0 {
		log.Println("分类已存在")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "分类已存在",
		})
		return
	}

	//创建分类
	category := model.Category{
		Identity: identity,
		Name:     name,
		ParentID: parentId,
	}

	err = define.DB.Model(&model.Category{}).Create(&category).Error
	if err != nil {
		log.Println("创建分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "创建分类失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建分类成功!",
	})
}

// DelCategory 删除分类
// @Tags 管理员方法
// @Summary 删除分类
// @Param token header string true "token"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /admin/delCategory [delete]
func DelCategory(c *gin.Context) {
	identity := c.Query("identity")
	//判断分类是否存在，或者有对应的问题关系
	var cnt int64
	err := define.DB.Model(&model.CategoryProblem{}).Where("category_id = (Select id from gorm_eo_category where identity = ? limit 1)", identity).Count(&cnt).Error
	if err != nil {
		log.Println("查找分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查找分类失败" + err.Error(),
		})
		return
	}
	log.Println(cnt)
	if cnt > 0 {
		log.Println("该分类下面已存在问题，不可删除")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "该分类下面已存在问题，不可删除",
		})
		return
	}

	err = define.DB.Where("identity = ?", identity).Delete(&model.Category{}).Error
	if err != nil {
		log.Println("删除分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "删除分类失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除分类成功!",
	})
}

// ModifyCategory 修改分类
// @Tags 管理员方法
// @Summary 修改分类
// @Param token header string true "token"
// @Param identity formData string true "identity"
// @Param name formData string true "name"
// @Param parent_id formData string true "parent_id"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /admin/modifyCategory [put]
func ModifyCategory(c *gin.Context) {
	identity := c.PostForm("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))

	//判断分类是否存在
	var cnt int64
	err := define.DB.Model(&model.Category{}).Where("identity = ?", identity).Count(&cnt).Error
	if err != nil {
		log.Println("查找分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查找分类失败" + err.Error(),
		})
		return
	}

	if cnt <= 0 {
		log.Println("分类不存在")
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "分类不存在",
		})
		return
	}

	//创建对象
	category := model.Category{
		Identity: identity,
		Name:     name,
		ParentID: parentId,
	}

	//更新数据
	err = define.DB.Model(&model.Category{}).Where("identity = ?", identity).Updates(&category).Error
	if err != nil {
		log.Println("更新分类失败", err)
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "更新分类失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
