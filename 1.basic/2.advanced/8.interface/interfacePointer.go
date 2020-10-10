package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}

func (d *Dog) Talk() {
	fmt.Printf("汪汪汪")
}
func (d *Dog) Eat() {
	fmt.Println("在吃骨头")
}
func (d *Dog) Name() string {
	fmt.Println("我是旺财")
	return "旺财"
}

func main() {
	var a Animal
	//var d Dog
	//a=d
	//a.Eat()

	fmt.Printf("%T %v\n", a, a)

	var d1 *Dog = &Dog{}
	a = d1

	a.Eat()
	fmt.Printf("*Dog %T %v\n", a, a)
}
