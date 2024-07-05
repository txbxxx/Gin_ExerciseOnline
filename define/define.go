/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 14:28
 * @File:  define
 * @Software: GoLand
 **/

package define

import "GinProject_ExerciseOnline/utils"

var (
	DefaultPage = "1"  // 默认显示页数
	DefaultSize = "10" // 每页显示个数
	RDB         = utils.RedisDBUtil()
	DB, _       = utils.DBUntil()
)
