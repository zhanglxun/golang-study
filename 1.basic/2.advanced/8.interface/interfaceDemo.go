package main

import "fmt"

type AnimalInterface interface {
	Talk1()
	Eat1()
	Name1() string
}

type DogStruct struct {
}

func (d *DogStruct) Talk1() {
	fmt.Println("汪汪~")
}
func (d *DogStruct) Eat1() {
	fmt.Println("吃骨头在")
}
func (d *DogStruct) Name1() string {
	fmt.Println("哆哆")
	return "哆哆"
}

//猪的定义
type Pig struct {
}

func (d *Pig) Talk1() {
	fmt.Println("哄哄~")
}
func (d *Pig) Eat1() {
	fmt.Println("吃草在")
}
func (d *Pig) Name1() string {
	fmt.Println("小猪-佩奇")
	return "小猪佩奇"
}

func main() {

	var ds DogStruct
	var aa AnimalInterface = &ds

	aa.Name1()
	aa.Talk1()
	aa.Eat1()

	var pig Pig

	aa = &pig

	aa.Name1()
	aa.Talk1()
	aa.Eat1()

}
