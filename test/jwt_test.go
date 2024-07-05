/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/5 14:20
 * @File:  jwt_test
 * @Software: GoLand
 **/

package test

import (
	"GinProject_ExerciseOnline/utils"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetMd5(t *testing.T) {
	md5 := utils.GetMd5("123456")
	fmt.Println(md5)
}

func TestGenerateToken(t *testing.T) {
	token, err := utils.GenerateToken("User_1", "张三", 1)
	if err != nil {
		fmt.Println("error：", err)
	}
	fmt.Println(token)
}

func TestAnalyseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6IlVzZXJfMSIsInBhc3N3b3JkIjoiIiwibmFtZSI6IuW8oOS4iSIsImlzX2FkbWluIjoxfQ.B8dT1e5W8rpWLEryJcpvG1Gp_HfSQS-L0BzbsTaedYQ"
	claims, err := utils.AnalyseToken(token)
	if err != nil {
		t.Fatal("error:", err) // 使用 t.Fatal 代替 fmt.Println
	}
	t.Logf("Claims: %v", claims) // 使用 t.Logf 代替 fmt.Println
}

func TestGenerateCode(t *testing.T) {
	//时间戳作为随机数种子
	rand.NewSource(time.Now().UnixNano())
	var code string
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%x", rand.Intn(10))
	}
	//生成随机数
	fmt.Println(code)
}
