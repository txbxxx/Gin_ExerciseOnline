/**
 * @Author tanchang
 * @Description 问题表
 * @Date 2024/7/4 0:24
 * @File:  Problem
 * @Software: GoLand
 **/

package model

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity;type:varchar(36)" json:"identity"`   //问题的唯一标识
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`        //问题标题
	Context    string `gorm:"column:context;type:text;" json:"context"`           //问题内容
	MaxMem     int    `gorm:"column:max_mem;type:int(11)" json:"max_mem"`         //最大运行内存
	MaxRuntime int    `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"` //最大运行时常
}
