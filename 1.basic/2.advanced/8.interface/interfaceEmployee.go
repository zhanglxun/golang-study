package main

import "fmt"

/*
 雇员，程序员，销售的薪资接口类例子
*/

//接口定义
type Employer interface {
	CalcSalary() float32
}

//开发人员的定义
type Programer struct {
	name  string
	base  float32
	extra float32
}

//开发人员构造函数
func NewProgramer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}
func (p Programer) CalcSalary() float32 {
	return p.base
}

//销售人员定义
type Sale struct {
	name  string
	base  float32
	extra float32
}

//销售人员构造函数
func NewSale(name string, base float32, extra float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}
func (p Sale) CalcSalary() float32 {
	return p.base + p.extra*p.base*0.5
}

//q求所有的工资
func calcAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost = cost + v.CalcSalary()
	}
	return cost
}

func main() {
	p1 := NewProgramer("jerry", 15000, 0)
	p2 := NewProgramer("tina", 12000, 0)
	p3 := NewProgramer("chris", 13000, 0)

	s1 := NewSale("lucy", 13000, 2.0)
	s2 := NewSale("marry", 16000, 2.1)
	s3 := NewSale("tom", 12000, 2.3)

	var employerList []Employer
	employerList = append(employerList, p1)
	employerList = append(employerList, p2)
	employerList = append(employerList, p3)

	employerList = append(employerList, s1)
	employerList = append(employerList, s2)
	employerList = append(employerList, s3)

	cost := calcAll(employerList)
	fmt.Printf("本月员工成本：%f \n", cost)
}
