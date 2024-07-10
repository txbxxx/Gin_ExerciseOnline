/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/4 22:17
 * @File:  submitSvc
 * @Software: GoLand
 **/

package service

import (
	"GinProject_ExerciseOnline/dao"
	"GinProject_ExerciseOnline/define"
	"GinProject_ExerciseOnline/model"
	"GinProject_ExerciseOnline/utils"
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// SearchSubmitList 提交列表
// @Tags 公共方法
// @Summary 查询提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param user_identity query string false "user_identity"
// @Param problem_identity query string false "problem_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /submit/searchSubmitList [get]
func SearchSubmitList(c *gin.Context) {
	//拿到默认显示页数、每页显示个数、用户唯一标识，问题唯一标识、提交状态
	page, _ := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	userIdentity := c.Query("user_identity")
	problemIdentity := c.Query("problem_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	//offset从0开始，所以，需要处理一下page
	page = (page - 1) * size

	//获取数据
	var count int64
	tx := dao.GetSubmitList(userIdentity, problemIdentity, status)
	submitList := make([]*model.Submit, 1)
	err := tx.Count(&count).Offset(page).Limit(size).Find(&submitList).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "查询失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"page":        page,
			"size":        size,
			"count":       count,
			"submit_list": submitList,
		},
	})
}

// Submit 提交代码
// @Tags 用户方法
// @Summary 提交代码
// @Param token header string true "token"
// @Param problem_identity query string false "problem_identity"
// @Param code body string true "code"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /user/submit [post]
func Submit(c *gin.Context) {
	identity := c.Query("problem_identity")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "获取代码错误" + err.Error(),
		})
		return
	}
	//将代码保存到服务器本地
	path, err := utils.SaveCode(body)
	if err != nil {
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "保存代码错误" + err.Error(),
		})
		return
	}

	//创建提交对象
	u, _ := c.Get("user_claims")
	user := u.(*utils.UserClaims)
	submit := &model.Submit{
		Identity:        utils.GenerateUUID(),
		ProblemIdentity: identity,
		UserIdentity:    user.Identity,
		Path:            path,
	}

	//通过问题的identity获取问题的测试案列
	problem := new(model.Problem)
	err = define.DB.Where("identity = ?", identity).Preload("TestCase").Take(&problem).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "获取题目错误" + err.Error(),
		})
		return
	}

	//使用协程来运行测试代码，并且通过不同管道来判断是否运行成功
	WA := make(chan int)  //答案错误的管道
	OOM := make(chan int) //超内存的管道
	CE := make(chan int)  //编译错误的管道
	AC := make(chan int)  //答案正确的管道
	//EC := make(chan int)  //非法代码的管道

	//通过的测试案列
	passCount := 0
	//提示信息
	var msg string
	//互斥锁
	var lock sync.Mutex

	//拿到测试案列并运行
	for _, testCase := range problem.TestCase {
		testCase := testCase
		fmt.Println("开始运行代码......")
		go func() {
			fmt.Println("代码存在: " + path)
			cmd := exec.Command("go", "run", path)
			var out, stderr bytes.Buffer
			cmd.Stderr = &stderr
			cmd.Stdout = &out
			stdinPipe, err := cmd.StdinPipe()
			if err != nil {
				log.Fatalln(err)
			}
			io.WriteString(stdinPipe, testCase.Input+"\n")

			//获取运行前的内存快照
			var beforeMem runtime.MemStats
			runtime.ReadMemStats(&beforeMem)

			//运行代码
			if err := cmd.Run(); err != nil {
				log.Println("error", stderr.String())
				//判断是否是编译错误
				if err.Error() == "exit status 1" {
					msg = stderr.String()
					CE <- 1
					return
				}
			}

			//获取运行后的内存快照
			var afterMem runtime.MemStats
			runtime.ReadMemStats(&afterMem)

			//判断是否是答案错误
			if out.String() != testCase.Output {
				msg = "答案错误"
				WA <- 1
				return
			}

			//判断是否运行超内存,转换成MB
			if (afterMem.Alloc-beforeMem.Alloc)/(1024*1024) > uint64(problem.MaxMem) {
				OOM <- 1
				return
			}

			lock.Lock()
			passCount++
			if passCount == len(problem.TestCase) {
				AC <- 1
			}
			lock.Unlock()
		}()
	}

	//根据管道信息来判断运行状态
	select {
	case <-WA:
		msg = "答案错误"
		submit.Status = 2
	case <-OOM:
		msg = "运行超内存"
		submit.Status = 4
	case <-CE:
		submit.Status = 5
	case <-AC:
		msg = "答案正确"
		submit.Status = 1
	case <-time.After(time.Millisecond * time.Duration(problem.MaxRuntime)):
		if passCount == len(problem.TestCase) {
			submit.Status = 1
			msg = "答案正确"
		} else {
			submit.Status = 3
			msg = "运行超时"
		}
	}

	//开启一个事务来运行sql
	err = define.DB.Transaction(
		func(tx *gorm.DB) error {
			err = tx.Create(submit).Error
			if err != nil {
				return errors.New("创建提交失败: " + err.Error())
			}
			//创建提交成功后，更新用户的提交和完成题目数，成功一次+1
			m := make(map[string]interface{})
			//提交一次+1
			m["submit_problem_num"] = gorm.Expr("submit_problem_num + ?", 1)
			//如果提交状态为成功成功次数就+1
			if submit.Status == 1 {
				m["finish_problem_num"] = gorm.Expr("finish_problem_num + ?", 1)
			}
			//更新当前用户的提交和完成题目数
			err = tx.Model(&model.User{}).Where("identity = ?", user.Identity).Updates(m).Error
			if err != nil {
				return errors.New("更新用户提交和完成题目数失败: " + err.Error())
			}

			//更新问题的提交次数和完成次数
			err = tx.Model(&model.Problem{}).Where("identity = ?", identity).Updates(m).Error
			if err != nil {
				return errors.New("更新问题提交和完成次数失败: " + err.Error())
			}
			return nil
		},
	)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "提交失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"status": submit.Status,
			"msg":    msg,
		},
	})
}
