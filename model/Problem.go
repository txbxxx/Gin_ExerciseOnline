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
	Identity         string     `gorm:"column:identity;type:varchar(36);unique" json:"identity"`  //问题的唯一标识
	Title            string     `gorm:"column:title;type:varchar(255)" json:"title"`              //问题标题
	Context          string     `gorm:"column:context;type:text;" json:"context"`                 //问题内容
	MaxMem           int        `gorm:"column:max_mem;type:int(11)" json:"max_mem"`               //最大运行内存【MB】
	MaxRuntime       int        `gorm:"column:max_runtime;type:int(11)" json:"max_runtime"`       //最大运行时常【毫秒】
	FinishProblemNum int64      `gorm:"column:finish_problem_num;type:int(11)" json:"pass_num"`   //问题通过次数
	SubmitProblemNum int64      `gorm:"column:submit_problem_num;type:int(11)" json:"submit_num"` //问题提交次数
	Category         []Category `gorm:"many2many:category_problem"`                               //连接关联表，表示问题和分类多对多关系
	TestCase         []TestCase `gorm:"foreignKey:ProblemIdentity;reference:Identity"`            //连接关联表，表示问题和测试用例一对多关系
}
