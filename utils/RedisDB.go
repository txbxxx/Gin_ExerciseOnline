/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/5 17:20
 * @File:  RedisDB
 * @Software: GoLand
 **/

package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisDBUtil() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RDBAddr,
		Password: RDBPwd,
		DB:       RDBDefaultDB,
	})
}
