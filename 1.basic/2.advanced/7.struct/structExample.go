package main

import "fmt"

/**
内部的一个写法
*/
type User struct {
	userName string
	sex      string
	age      int
	url      string
}

// 匿名字段玩法
type Student struct {
	name string
	sex  string
	int
	string
}

//go 里面没有构造函数，需自己定义
func NewUser(userName string, sex string, age int, url string) *User {
	user := &User{
		userName: userName,
		sex:      sex,
		age:      age,
		url:      url,
	}

	// 第二种赋值方式
	//user:=new(User)
	//user.age=18
	//user.url="google.com"
	//user.sex="nan"
	//user.userName="jarvis"

	return user
}

func main() {

	//直接给结构体赋值
	user := User{
		userName: "jarvis",
		sex:      "男",
		age:      30,
		url:      "google.com",
	}

	fmt.Printf("users = %v \n", user)

	//定义后赋值
	var user2 User
	user2.age = 20
	user2.url = "qq.com"
	user2.userName = "jarvis2chang"
	user2.sex = "男"

	fmt.Printf("user2 = %v \n", user2)

	// 定义后，在赋值
	var user9 User
	user9.age = 18
	user9.url = "google.com"
	user9.sex = "nan"
	user9.userName = "jarvis"

	fmt.Printf("username= %s,age =%d,sex=%s,url=%s \n", user9.userName, user9.age, user9.sex, user9.url)

	//匿名字段的玩法，默认使用类型名字 做字段名字
	stu := &Student{
		name:   "abc",
		sex:    "nv",
		int:    23,
		string: "url",
	}
	fmt.Printf("stu=%#v \n", stu)

	/**
	结构体的玩法
	*/

	//直接给结构体赋值
	user8 := User{
		userName: "lixun",
		sex:      "nv",
		age:      33,
		url:      "youtube.com",
	}
	fmt.Printf("user2=%#v \n", user8)
	//为 nil 的空对象
	var user3 User
	fmt.Printf("user3=%#v \n", user3)

	//指针 定义 结构体，指针需要初始化，才能使用,结构体的指针 简化了 *user5，直接user5
	var user4 *User
	fmt.Printf("user4=%#v \n", user4)

	/*
	 指针结构体，初始化的方式两种
	 1、&User
	 2、new(User)
	*/

	// &User 分配了一个User类型的对象，并返回一个指针，&User 初始化
	var user5 = &User{}
	user5.age = 19
	user5.sex = "nan"
	user5.userName = "jerry"
	user5.url = "bing.com"
	fmt.Printf("user5=%#v \n", user5)

	var user6 = &User{
		userName: "lixun",
		sex:      "nv",
		age:      33,
		url:      "youtube.com",
	}
	fmt.Printf("user6=%#v \n", user6)

	// new(User) 初始化
	var user7 *User = new(User)
	user7.age = 19
	user7.sex = "nan"
	user7.userName = "jerry"
	user7.url = "bing.com"
	fmt.Printf("user7=%#v \n", user7)

	//构造函数来弄
	user10 := NewUser("user01", "nv", 19, "gouzaohanshu")
	fmt.Printf("user8=%#v \n", user10)
}
