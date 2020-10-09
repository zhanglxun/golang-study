package main

import (
	"fmt"
	"reflect"
)

//反射的基本练习
func testReflect1() {

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
}

/*
 使用一个接口传递参数
反射拿到类型
*/
func testReflect2(a interface{}) {

	t := reflect.TypeOf(a)
	fmt.Printf("type of a is %v \n", t)

	k := t.Kind()
	switch k {
	case reflect.Int32:
		fmt.Printf("type is int32\n")
	case reflect.String:
		fmt.Printf("type is string\n")
	case reflect.Int64:
		fmt.Printf("type is int64\n")
	case reflect.Float32:
		fmt.Printf("type is float32\n")
	case reflect.Float64:
		fmt.Printf("type is float64\n")
	default:
		fmt.Printf("can't find type,%v\n", k)
	}
}

func testReflect3(b interface{}) {

	//t:=reflect.TypeOf(b) 这两写法都可以
	v := reflect.ValueOf(b)

	t := v.Kind()
	switch t {
	case reflect.Int32:
		fmt.Printf("type is int32,value is %d\n", v.Int())
	case reflect.String:
		fmt.Printf("type is string,value is %s\n", v.String())
	case reflect.Int64:
		fmt.Printf("type is int64,value is %d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32,value is %f\n", v.Float())
	case reflect.Float64:
		fmt.Printf("type is float64,value is %f\n", v.Float())
	default:
		fmt.Printf("can't find type,%v\n", t)
	}
}

// 通过接口，指针类型，来赋值
func testRelfectValue(a interface{}) {
	v := reflect.ValueOf(a)

	k := v.Kind()
	switch k {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("type is int64,value is %d\n", v.Int())
	case reflect.Float64:
		v.SetFloat(100)
		fmt.Printf("type is float64,value is %f\n", v.Float())
	// ptr 来判断 指针，通过elem()来 设置值
	case reflect.Ptr:
		fmt.Printf("tset value to int \n")
		v.Elem().SetInt(67)

	default:
		fmt.Printf("default ")

	}

	fmt.Printf("v type is %T ,value is %v\n", v.Type(), v)
}

func main() {

	testReflect1()

	var x int64 = 34
	testReflect2(x)

	var r = "jarvis"
	testReflect3(r)

	fmt.Println("------")

	var c int64 = 45
	/* 如果直接传 c 的值，会reflect.Value.SetInt using unaddressable value
	   int 的值类型，直接传值是会报错的，只能传地址 &c
	   通过副本卖给 值类型赋值，会报错
	*/
	testRelfectValue(&c)
	fmt.Printf("c value is %v \n", c)
}
