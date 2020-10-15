package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
第一个简单使用 gorm 的例子，非常简单的一个入门的例子
*/

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {

	db, err := gorm.Open("mysql", "pxndev:pxndev@2018@(rm-wz97vpexguxsff7en5o.mysql.rds.aliyuncs.com:3306)/db_apiserver?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "七米", "男", "篮球"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)

}
