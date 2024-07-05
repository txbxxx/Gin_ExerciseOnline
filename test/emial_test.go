/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/5 15:39
 * @File:  emial_test
 * @Software: GoLand
 **/

package test

import (
	"GinProject_ExerciseOnline/utils"
	"testing"
)

func TestSendEmail(t *testing.T) {
	err := utils.SendEmail("", "123456")
	if err != nil {
		t.Error(err)
	}
}
