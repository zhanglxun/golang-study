package main

import (
	"encoding/json"
	"fmt"
)

/*
 json 序列化
*/

type Students struct {
	Id   int
	Name string
	Sex  string
}

type ClsssDemo struct {
	Name     string
	Count    int
	Students []*Students
}

var strJson = `{"Name":"一年级","Count":0,"Students":[{"Id":0,"Name":"张三","Sex":"男"},{"Id":1,"Name":"张三","Sex":"男"},{"Id":2,"Name":"张三","Sex":"男"},{"Id":3,"Name":"张三","Sex":"男"},{"Id":4,"Name":"张三","Sex":"男"},{"Id":5,"Name":"张三","Sex":"男"},{"Id":6,"Name":"张三","Sex":"男"},{"Id":7,"Name":"张三","Sex":"男"},{"Id":8,"Name":"张三","Sex":"男"},{"Id":9,"Name":"张三","Sex":"男"}]}`

/*
 json 反序列化
*/
func serilaziable() {

	var c1 *ClsssDemo = &ClsssDemo{}

	err := json.Unmarshal([]byte(strJson), c1)

	if err != nil {
		fmt.Println("unmarshal json failed")
		return
	}
	fmt.Printf("c1 :%#v \n", c1)

	for _, v := range c1.Students {
		fmt.Printf("student:%#v \n", v)
	}
}

func main() {
	c := &ClsssDemo{
		Name:  "一年级",
		Count: 10,
	}
	for i := 0; i < 10; i++ {
		stu := &Students{
			Name: "张三",
			Id:   i,
			Sex:  "男",
		}
		c.Students = append(c.Students, stu)
	}
	// json 字符串输出
	data, err := json.Marshal(c)

	if err != nil {
		fmt.Println("json failed")
		return
	}

	fmt.Printf("json=%s \n", string(data))

	serilaziable()
}
