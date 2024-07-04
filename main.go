/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 0:59
 * @File:  main
 * @Software: GoLand
 **/

package main

import (
	"GinProject_ExerciseOnline/router"
)

func main() {
	//创建GinServer服务
	r := router.Router()
	//启动服务，默认是8080端口
	err := r.Run(":8082")
	if err != nil {
		return
	}
}
