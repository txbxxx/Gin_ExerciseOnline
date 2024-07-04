/**
 * @Author tanchang
 * @Description 提交表
 * @Date 2024/7/4 0:51
 * @File:  Submit
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36)" json:"identity"`                 // 提交的的唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"` // 提交的的唯一标识
	UserIdentity    string `gorm:"column:user_identity;type:varchar(36)" json:"user_identity"`       // 提交的的唯一标识
	Path            string `gorm:"column:path;type:varchar(255)" json:"path"`                        //提交代码存储路径
	Status          int    `gorm:"column:status;type:tinyint(1)" json:"status"`
}
