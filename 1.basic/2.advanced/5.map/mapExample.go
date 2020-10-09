package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

/*
 map 的相关操作
 1、map是一个应用类型 key-value 类型
*/
func testMap1() {
	var a map[string]int
	fmt.Printf("a=%v\n", a)
	if a == nil {

		//使用make初始化 map
		a = make(map[string]int, 16)
		fmt.Printf("a=%v \n", a)

		a["stu1"] = 1000
		a["stu2"] = 2000
		a["stu3"] = 3000
		fmt.Printf("a=%v \n", a)

	}

	//直接赋值
	b := map[string]int{
		"jar01": 100,
		"jar02": 301,
		"jar03": 409,
	}
	fmt.Printf("b=%v \n", b)

	//通过 key 取值 ，b[key] 的方式取值
	var key string = "jar03"
	fmt.Printf("value of key[%s] is %d \n", key, b[key])

	// 判断map中的key 是否存在，这是一个特定语法。【重要】
	result, ok := b["jar03"]

	if ok == false {
		fmt.Printf("key %s is not exist\n", "jar04")
	} else {
		fmt.Printf("key %s is %d \n", "jar04", result)
	}

}

// 操作map 中的元素,随机产生key-value 到map ，遍历取出
func testMap2() {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int, 10)

	fmt.Printf("a type is %T \n", a)

	for i := 0; i < 10; i++ {
		//字符串拼接
		key := fmt.Sprintf("jarvis%d", i)
		value := rand.Intn(10)

		a[key] = value
	}

	//for key,value:=range a{
	//	fmt.Printf("map[%s]=%d \n",key,value)
	//}

	//排序 操作,字符的 顺序
	var keys []string = make([]string, 0, 10)

	fmt.Printf("keys type is %T \n", keys)

	for key, value := range a {
		fmt.Printf("key[%s]=%d \n", key, value)
		keys = append(keys, key)
	}

	fmt.Println("keys value is  : ", keys)

	sort.Strings(keys)
	for _, value := range keys {
		fmt.Printf("key-sort[%s]=%d \n", value, a[value])
	}

}

//删除map中的元素 delete(T,key)
func testMap3() {
	//直接赋值
	c := map[string]int{
		"jar01": 100,
		"jar02": 301,
		"jar03": 409,
	}
	delete(c, "jar01")
	fmt.Printf("c=%v \n", c)

}

// map 的切片操作，map是引用类型，【重点】
func modifyMap4(a map[string]int) {
	a["j1"] = 100
}
func testmap4() {
	var a map[string]int = make(map[string]int, 10)
	a["j2"] = 30
	a["j3"] = 40
	modifyMap4(a)
	fmt.Printf("a=%v \n", a)
}

/* 1.slice里面装map
   2.map里面装slice
*/

//切片 里面放入 map，类似 于java 里面 list<HashMap<sting,int>> 这样的数据结构 【重要】
func testmap5() {
	var s []map[string]int
	//为切片初始化，但是里面的元素，并没有初始化
	s = make([]map[string]int, 5, 10)
	for index, value := range s {
		fmt.Printf("slice[%d]=%v \n", index, value)
	}

	//为map 进行初始化，分配内存
	s[0] = make(map[string]int, 10)
	s[0]["go1"] = 10
	s[0]["go2"] = 20
	s[0]["go3"] = 30

	for index, value := range s {
		fmt.Printf("slice after make s[%d]=%v \n", index, value)
	}
}

// map 里面装一个切片，类似java 里面 hashMap<string,list<T>> 这样的数据结构 【重要】
func testMap6() {
	rand.Seed(time.Now().UnixNano())
	var s map[string][]int
	s = make(map[string][]int, 10)
	//s["golang1"]=make([]int,5,10)
	key := "golang"
	value, ok := s[key]

	println(value, ok)

	if !ok {
		s[key] = make([]int, 2, 10)
		value = s[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	value = append(value, 300)
	s[key] = value
	fmt.Printf("map:%v \n", s)
}

/*
 1、写个程序，统计一个字符串，每个单词出现的次数，如 s="how do you do",输出，how=1，do=2,you=1
 2、写一个程序，实现学生信息存储，学生有id，年龄，分数等信息，需要方便的通过id，查询对应学生的信息
*/

func sumWorld() {
	var str string = "how do you do , do you have dinner ?"
	//定义一个切片，存储分词后的 slice
	sliceWorld := strings.Split(str, " ")
	//定义一个map，进行 分词的统计
	var worldCount map[string]int = make(map[string]int, 100)
	for _, v := range sliceWorld {
		//每次循环，先判断该单词是否存在，不存在则创建,并赋值为1
		fmt.Printf("v =%v \n", v)
		count, ok := worldCount[v]
		fmt.Printf("count=%v,ok=%v \n", count, ok)
		if !ok {
			worldCount[v] = 1
		} else {
			worldCount[v] = count + 1
		}
	}

	fmt.Printf("result=%v \n", worldCount)
}

// 学生成绩分数 的数据结构 【重要】
func studentStore() {
	//map 里面int 是id, 类似java 里面的数据结构 hashMap<int,map<string, interface>>
	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{}, 16)
	//插入学生信息 id=1，姓名的=loe，分数=78.5，年龄=19
	var id = 1
	var name = "leo"
	var age = 18
	var score = 78.5

	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
	}

	value["name"] = name
	value["id"] = id
	value["age"] = age
	value["score"] = score

	stuMap[id] = value

	fmt.Printf("stuMap=%v \n", stuMap)

	for i := 0; i < 10; i++ {
		value, ok := stuMap[i]
		if !ok {
			value = make(map[string]interface{}, 8)
		}

		value["name"] = fmt.Sprintf("jarvis%d", i)
		value["id"] = i
		value["age"] = rand.Intn(100)
		value["score"] = rand.Float32() * 100.0

		stuMap[i] = value
	}

	for k, v := range stuMap {

		fmt.Printf("id=%d, stuInfo=%#v \n", k, v)
	}
}

func main() {

	//testMap1()

	//随机产生key-value 到map ，遍历取出
	//testMap2()

	// 删除一个值
	//testMap3()

	//
	//testmap4()

	//testmap5()

	//testMap6()

	//sumWorld()

	studentStore()
}
