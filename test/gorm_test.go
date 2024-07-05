/**
 * @Author tanchang
 * @Description //数据库测试
 * @Date 2024/7/4 4:06
 * @File:  gorm_test
 * @Software: GoLand
 **/

package test

import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	//log.Println(define.DB.Name())
	//category1 := model.Category{
	//	Identity: "1",
	//	Name:     "算法",
	//	ParentID: 0,
	//	Problem: []model.Problem{
	//		{
	//			Identity:   "1",
	//			Title:      "斐波那契",
	//			Context:    "斐波那契",
	//			MaxMem:     1,
	//			MaxRuntime: 1,
	//		},
	//		{
	//			Identity:   "2",
	//			Title:      "塔罗牌",
	//			Context:    "塔罗牌",
	//			MaxMem:     2,
	//			MaxRuntime: 2,
	//		},
	//	},
	//}

	//problem1 := model.Problem{
	//	Identity:   "3",
	//	Title:      "冒泡排序",
	//	Context:    "冒泡排序",
	//	MaxMem:     3,
	//	MaxRuntime: 3,
	//	Category: []model.Category{
	//		//category1,
	//		{
	//			Identity: "2",
	//			Name:     "循环",
	//			ParentID: 0,
	//		},
	//	},
	//}

	var p model.Problem
	tx := define.DB.Debug().
		Preload("Category").
		Joins("LEFT JOIN gorm_eo_category_problem on gorm_eo_category_problem.problem_id = gorm_eo_problem.id").
		Where("gorm_eo_category_problem.category_id = (select gorm_eo_category.id from gorm_eo_category where gorm_eo_category.identity = ? )", 1).Find(&p)
	if tx.Error != nil {
		t.Error("创建类别错误", tx.Error)
	}

	fmt.Println(p)
	//define.DB.Model(&model.Problem{}).Create(&problem)

}
