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
	"testing"
)

func TestName(t *testing.T) {
	proble := []model.Problem{
		{
			Identity:   "1",
			CategoryId: "1",
			Title:      "斐波那契",
			Context:    "斐波那契",
			MaxMem:     1,
			MaxRuntime: 1,
		},
		{
			Identity:   "2",
			CategoryId: "2",
			Title:      "塔罗牌",
			Context:    "塔罗牌",
			MaxMem:     2,
			MaxRuntime: 2,
		},
		{
			Identity:   "3",
			CategoryId: "3",
			Title:      "冒泡排序",
			Context:    "冒泡排序",
			MaxMem:     3,
			MaxRuntime: 3,
		},
	}

	tx := define.DB.Model(&model.Problem{}).Create(&proble)
	if tx.Error != nil {
		t.Error(tx.Error)
	}
}
