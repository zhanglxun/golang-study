package main

import (
	"encoding/json"
	"fmt"
)

type UserBase struct {
	UserName string `json:"user_name"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Url      string
	Address  Address
}

type Address struct {
	Province string
	City     string
}

type StudentBase struct {
	Name       string
	Age        int
	CreateTime string
	City       string
	*Address
}

/*

 */

//结构体的嵌套 使用
func studentTest1() {
	userDemo := UserBase{
		UserName: "章文",
		Sex:      "女",
		Url:      "qq.com",
		Age:      19,
		Address: Address{
			Province: "",
			City:     "",
		},
	}

	fmt.Printf("UserBase:%#v \n", userDemo)
}

// 结构体冲突，的写法
func UserBaseFunc() {
	var student StudentBase
	student.Name = ""
	student.Age = 19
	student.Address = new(Address)

	student.Address.City = "深圳"
	fmt.Printf("student:%#v ,address-city:%#v \n", student, student.Address.City)

}

// struts里面的tag，用反射取值,这个很有用，后面的json 返回值，都可能用到 【重要】
// 结构体 中 字段需要大写，json 包 才能访问到，否则私有变量，输出json  为空
func structTest5() {
	UserBase := &UserBase{
		UserName: "jarviszhang",
		Sex:      "男",
		Age:      19,
		Url:      "abc.com",
	}
	// 反射取值，json 化
	date, _ := json.Marshal(UserBase)
	fmt.Printf("json str :%s \n", string(date))
}

func main() {

	//struct 嵌套用法
	studentTest1()

	UserBaseFunc()

	structTest5()
}
