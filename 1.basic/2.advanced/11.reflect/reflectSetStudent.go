package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
	//xxx   int
}

func main() {

	var a = 1
	b := reflect.ValueOf(a)

	var s Student
	v := reflect.ValueOf(&s)

	//通过反射拿到 元素，给字段赋值
	v.Elem().Field(0).SetString("stu01")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().FieldByName("Age").SetInt(18)
	v.Elem().FieldByName("Score").SetFloat(99.0)

	fmt.Printf("b value :%v \n", b)
	fmt.Printf("s value :%v \n", v)

}
