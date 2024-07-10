/**
 * @Author tanchang
 * @Description //数据库连接设置
 * @Date 2024/7/5 17:53
 * @File:  DbSetting
 * @Software: GoLand
 **/

package utils

var (
	DBAddr       = "localhost:3306"    // 数据库地址
	DBUser       = "root"              // 数据库用户名
	DBPwd        = "000000"            // 数据库密码
	RDBAddr      = "172.20.241.3:6379" // Redis地址
	RDBPwd       = ""                  // Redis密码
	RDBDefaultDB = 0                   // Redis默认数据库
)
