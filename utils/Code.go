/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/9 16:07
 * @File:  Code
 * @Software: GoLand
 **/

package utils

import "os"

// SaveCode 保存用户运行的代码
func SaveCode(code []byte) (string, error) {
	dirName := "code/" + GenerateUUID()
	filePath := dirName + "/main.go"
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return "", err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	file.Write(code)
	defer file.Close()
	return filePath, nil
}
