package main

import "fmt"

type People struct {
	Name    string
	Country string
}

func (p People) PrintPeople() {
	fmt.Printf("name=%s,country=%s \n", p.Name, p.Country)
}

// 值 类型 和指针类型的区别，值对象较大时候，在拷贝时的开销，可能较大，建议用指针类型
func (p People) Set(name string, country string) {
	p.Name = name
	p.Country = country
}

/*
 1、需要修改接受者中的值
 2、接受者是 大对象，副本拷贝代价大
 3、通常使用指针类型作为接收者
*/

// struct 是值类型，修改传参 需要用指针
func (p *People) SetPoint(name string, country string) {
	p.Name = name
	p.Country = country
}

func main() {

	p1 := People{
		Name:    "jarvis",
		Country: "china",
	}
	p1.PrintPeople()

	p1.Set("jerry", "england")
	p1.PrintPeople()

	// 显示的 (&p1).setPoint() ，语法糖
	p1.SetPoint("lucy", "jpa")
	p1.PrintPeople()
}
