/**
 * @Author tanchang
 * @Description 数据库工具方法
 * @Date 2024/7/4 4:24
 * @File:  DBconn
 * @Software: GoLand
 **/

package utils

import (
	"GinProject_ExerciseOnline/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

// DBUntil 用于连接数据库
func DBUntil() (*gorm.DB, error) {
	databases := "root:000000@tcp(localhost:3306)/EO?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	db, err := gorm.Open(mysql.Open(databases), &gorm.Config{
		SkipDefaultTransaction: false, //禁用事务
		NamingStrategy: schema.NamingStrategy{ //命名策略
			TablePrefix:   "gorm_eo_", //gorm创建表添加此前缀
			SingularTable: true,       //禁用复数名称
		},
	})
	if err != nil {
		log.Println("数据库连接失败", err.Error())
		return nil, err
	}

	err = db.SetupJoinTable(&model.Problem{}, "Category", &model.CategoryProblem{})
	if err != nil {
		log.Println("设置JoinTable失败", err.Error())
		return nil, err
	}

	err = db.SetupJoinTable(&model.Category{}, "Problem", &model.CategoryProblem{})
	if err != nil {
		log.Println("设置JoinTable失败", err.Error())
		return nil, err
	}

	//如果数据表已经存在就不在自动迁移，如果不存在则自动迁移(也就是创建数据表)
	if !(db.Migrator().HasTable(&model.User{}) &&
		db.Migrator().HasTable(&model.Category{}) &&
		db.Migrator().HasTable(&model.Problem{}) &&
		db.Migrator().HasTable(&model.Submit{}) &&
		db.Migrator().HasTable(&model.CategoryProblem{})) {
		db.AutoMigrate(&model.User{}, &model.Category{}, &model.Problem{}, &model.Submit{}, &model.CategoryProblem{})
	} else {
		log.Println("数据库表已存在")
	}
	return db, nil
}
