/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 19:28
 * @File:  ProblemDao
 * @Software: GoLand
 **/

package dao

import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"gorm.io/gorm"
)

// GetProblemList 向数据库中获取问题列表
func GetProblemList(searchData, categoryIdentity string) *gorm.DB {
	tx := define.DB.Model(&model.Problem{}).Preload("Category").Where("title like ? OR context like ?", searchData, searchData)
	if categoryIdentity != "" {
		tx = tx.Model(&model.Problem{}).
			Joins("LEFT JOIN gorm_eo_category_problem on gorm_eo_category_problem.problem_id = gorm_eo_problem.id").
			Where("gorm_eo_category_problem.category_id = (select gorm_eo_category.id from gorm_eo_category where gorm_eo_category.identity = ? )", categoryIdentity)
	}
	return tx
}
