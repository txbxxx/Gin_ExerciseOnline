/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/8 14:37
 * @File:  TestCase
 * @Software: GoLand
 **/

package model

import (
	"gorm.io/gorm"
)

type TestCase struct {
	gorm.Model
	Identity        string `gorm:"column:identity;type:varchar(36);unique" json:"identity"`          // 测试案列唯一标识
	ProblemIdentity string `gorm:"column:problem_identity;type:varchar(36)" json:"problem_identity"` //	问题的唯一表标识
	Input           string `gorm:"column:input;type:text" json:"input"`                              //输入案列
	Output          string `gorm:"column:output;type:text" json:"output"`                            //输出案例
}
