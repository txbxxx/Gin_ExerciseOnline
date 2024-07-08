/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/5 20:46
 * @File:  redis_test
 * @Software: GoLand
 **/

package test

import (
	"GinProject_ExerciseOnline/define"
	"context"
	"fmt"
	"testing"
	"time"
)

var ctx = context.Background()

func TestRDB(t *testing.T) {
	define.RDB.Set(ctx, "test", "test", time.Second*5)
	result, err := define.RDB.Get(ctx, "test").Result()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
