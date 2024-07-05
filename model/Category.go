/**
 * @Author tanchang
 * @Description 问题分类表
 * @Date 2024/7/4 0:49
 * @File:  Category
 * @Software: GoLand
 **/

package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string    `gorm:"column:identity;type:varchar(36);unique" json:"identity"` // 问题分类的唯一标识
	Name     string    `gorm:"column:name;type:varchar(100)" json:"name"`               // 问题分类名
	ParentID int       `gorm:"column:parent_id;type:int(11)" json:"parent_id"`          // 父级ID号
	Problem  []Problem `gorm:"many2many:category_problem"`                              // 连接关联表，表示问题和分类多对多关系
}
