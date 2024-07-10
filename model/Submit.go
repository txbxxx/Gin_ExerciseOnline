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
	Identity        string  `gorm:"column:identity;type:varchar(36);unique" json:"identity"`          // 提交的的唯一标识
	ProblemIdentity string  `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"` // 提交的的唯一标识
	Problem         Problem `gorm:"foreignKey:ProblemIdentity;references:Identity"`                   // 对应的题目
	UserIdentity    string  `gorm:"column:user_identity;type:varchar(36)" json:"user_identity"`       // 提交的的唯一标识
	User            User    `gorm:"foreignKey:UserIdentity;references:Identity"`                      // 对应的用户
	Path            string  `gorm:"column:path;type:varchar(255)" json:"path"`                        //提交代码存储路径
	Status          int     `gorm:"column:status;type:tinyint(1);default:1" json:"status"`            //代码的状态  1表示待判断，2表示正确，3表示答案错误，4表示运行超时
}
