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
	Identity string `gorm:"column:identity;type:varchar(36);unique" json:"identity"` // 用户的唯一标识
	Name     string `gorm:"column:name;type:varchar(100)" json:"name"`               // 用户名
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`        // 用户密码
	Phone    string `gorm:"column:phone;varchar(20)" json:"phone"`                   // 电话
	Mail     string `gorm:"column:mail;varchar(100)" json:"mail"`                    // 电子邮件
}
