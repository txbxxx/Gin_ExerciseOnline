/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 22:32
 * @File:  SubmitDao
 * @Software: GoLand
 **/

package dao

import (
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"gorm.io/gorm"
)

// GetSubmitList 列出提交列表
func GetSubmitList(userIdentity, problemIdentity string, status int) *gorm.DB {
	//过滤掉content和password
	tx := define.DB.Model(&model.Submit{}).Preload("Problem",
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("context")
		},
	).Preload("User",
		func(db *gorm.DB) *gorm.DB {
			return db.Omit("password")
		},
	)
	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}
	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}
	if status != 0 {
		tx.Where("status = ?", status)
	}

	return tx
}
