package main

import "fmt"

/*
 1、 go 里面，是通过匿名字段来实现集继承
 2、匿名字段需要使用指针
*/

type Animal struct {
	AnimalType string
	sex        string
}

func (animal *Animal) Talk() {
	fmt.Printf("animal talk : my type is %s \n", animal.AnimalType)
}

type Dog struct {
	Feet int
	run  bool
	*Animal
}

func (d *Dog) Eat() {
	fmt.Printf("dog is eat \n")
}

func main() {
	var d *Dog = &Dog{
		Feet: 4,
		run:  true,
		Animal: &Animal{
			AnimalType: "dog",
			sex:        "meal",
		},
	}

	d.Eat()
	d.Talk()
}
