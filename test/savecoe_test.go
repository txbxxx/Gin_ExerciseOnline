/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/10 9:55
 * @File:  savecoe_test
 * @Software: GoLand
 **/

package test

import (
	"GinProject_ExerciseOnline/utils"
	"os"
	"testing"
)

func TestSaveCode(t *testing.T) {
	dirName := "code/" + utils.GenerateUUID()
	filePath := dirName + "/main.go"
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
}
