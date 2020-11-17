package main

import "fmt"

func main() {
	var a = Animal{
		name: "dog",
		age:  1,
	}
	a.add1()
	fmt.Println(a)
	a.add2()
	fmt.Println(a)
}

type Animal struct {
	name string
	age  int
}

func (a Animal) add1() {
	a.age = a.age + 1
}

func (a *Animal) add2() {
	a.age = a.age + 1
}
