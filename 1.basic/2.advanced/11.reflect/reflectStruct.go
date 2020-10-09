package main

import (
	"fmt"
	"reflect"
)

type StudentRef struct {
	Name  string `json:"name"`
	Sex   int    `json:"sex"`
	Age   int    `json:"age"`
	Score float32
	Kid   int // 小写的字段数量是能获取到,但是反射拿值是拿不到的
}

func (s *StudentRef) SetName(name string) {
	s.Name = name
}
func (s *StudentRef) SetAge(age int) {
	s.Age = age
}

func (s *StudentRef) Print() {
	fmt.Printf("通过反射进行调用：%#v\n", s)
}

func testRef1() {
	var s StudentRef
	v := reflect.ValueOf(s)
	t := v.Type()
	//写法等价于
	//t:=reflect.Typeof(s)

	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Printf("s is int64 \n")
	case reflect.Struct:
		fmt.Printf("s is struct \n")
		fmt.Printf("s field number is %d \n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("field name %s, type is %v, value %v \n", t.Field(i).Name, t.Field(i).Type, field.Interface())
		}
	default:
		fmt.Printf("default")

	}
}

//通过反射 设置结构体的值
func testRef2() {
	var s StudentRef
	v := reflect.ValueOf(&s)

	v.Elem().Field(0).SetString("jarvis")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().FieldByName("Age").SetInt(33)
	v.Elem().FieldByName("Score").SetFloat(98.00)
	fmt.Printf("s:%#v \n", s)
}

// 指针类型传入 valueOf
func testRef3() {
	var s StudentRef
	//传s 则 是0 个方法，指针类型，就有1方法，要让值类型也有方法，那么在 方法的类型里面，就去掉指针。
	v := reflect.ValueOf(&s)
	t := v.Type()

	fmt.Printf("s has method number :%d \n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("struct %d method ,name :%s ,type:%v \n", i, method.Name, method.Type)
	}
}

//结构体方法，获取和调用
func testRef4() {
	//声明类,但是都还未赋值
	var s StudentRef
	//通过反射，获取Value 的信息
	v := reflect.ValueOf(&s)
	//t:=v.Type()

	// 根据方法名称，反射获取方法体
	m1 := v.MethodByName("Print")
	var args []reflect.Value
	m1.Call(args)

	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value
	name := "jarvisZhang"
	nameVal := reflect.ValueOf(name)

	//score:=98.5
	//scoreVal:=reflect.ValueOf(score)

	args2 = append(args2, nameVal)
	//args2=append(args2,scoreVal)

	m2.Call(args2)

	m1.Call(args)

}

// 结构体的tag 的信息获取，调用
func testRef5() {
	var s StudentRef
	//通过反射，获取Value 的信息
	v := reflect.ValueOf(&s)
	t := v.Type()

	field0 := t.Elem().Field(0)
	fmt.Printf("tag :json=%s \n", field0.Tag.Get("json"))
}
func main() {

	testRef1()

	testRef2()

	testRef3()

	testRef4()

	testRef5()
}
