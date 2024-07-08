/**
 * @Author tanchang
 * @Description 用户表
 * @Date 2024/7/4 0:41
 * @File:  User
 * @Software: GoLand
 **/

package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identity         string `gorm:"column:identity;type:varchar(36);unique" json:"identity"`                     // 用户的唯一标识
	Name             string `gorm:"column:name;type:varchar(100);unique" json:"name"`                            // 用户名
	Password         string `gorm:"column:password;type:varchar(32)" json:"password"`                            // 用户密码
	Phone            string `gorm:"column:phone;varchar(20)" json:"phone"`                                       // 电话
	Mail             string `gorm:"column:mail;varchar(100)" json:"mail"`                                        // 电子邮件
	IsAdmin          int    `gorm:"column:isadmin;type:tinyint(1);default:2" json:"is_admin"`                    // 是否为管理员1为是2为不是
	FinishProblemNum int    `gorm:"column:finish_problem_num;type:int(11);default:0;" json:"finish_problem_num"` // 问题完成个数
	SubmitProblemNum int    `gorm:"column:submit_problem_num;type:int(11);default:0;" json:"submit_problem_num"` // 提交文件次数
}
