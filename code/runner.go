/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/8 13:12
 * @File:  runner
 * @Software: GoLand
 **/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	//通过exec 执行代码
	cmd := exec.Command("go", "run", "code-user/main.go")
	var stdout, stderr bytes.Buffer

	//将标准输出和错误输出重定向到缓冲区
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	//创建标准输入管道，StdinPipe返回一个可读可写的管道，用于父进程向子进程发送数据
	pipe, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	//使用io.WriteString将需要输入的参数写入管道
	io.WriteString(pipe, "30 33\n")
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
	}

	//输出结果
	fmt.Println(stdout.String())
	fmt.Println(stderr.String())

	//判断测试结果
	if stdout.String() == "63" {
		fmt.Println("测试通过")
	} else {
		fmt.Println("测试失败")
	}
}
