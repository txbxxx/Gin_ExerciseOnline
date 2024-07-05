/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 15:37
 * @File:  CategoryProblem
 * @Software: GoLand
 **/

package model

import (
	"gorm.io/gorm"
	"time"
)

// CategoryProblem 问题和类别的关联表
type CategoryProblem struct {
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	CategoryID uint           `gorm:"column:category_id;type:int(11);primaryKey" json:"category_id"` // 关联表的分类ID
	ProblemID  uint           `gorm:"column:problem_id;type:int(11);primaryKey" json:"problem_id"`   // 关联表的问题ID
}
