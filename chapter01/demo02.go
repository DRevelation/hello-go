package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const PI = 10
const (
	A = 1
	B = 2
	C = 3
)

var str = ""
var (
	a = 1
	b = 2
	c = 3
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	fmt.Println("=========================")

	var b1 = []byte("abc")
	fmt.Println(b1)

	err := errors.New("this is error")
	if err != nil {
		fmt.Println(err)
	}

	var ages = make(map[string]int)
	ages["alice"] = 12
	ages["bob"] = 13
	for name, age := range ages {
		fmt.Printf("%s's age is %d\n", name, age)
	}

	var p = person{"candy", 14}
	fmt.Println(p.String())
	fmt.Println(p)
	modify(p)
	fmt.Println(p)
	modify2(&p)
	fmt.Println(p)

	//result := 10/0
	//fmt.Println(result)

	file, err := os.Open("D:\\gitee\\hello-go\\chapter01\\map01.go")
	if err != nil {
		log.Fatal("err=", err)
	}
	fmt.Println(file.Name())

	print("aa", "bb", "cc")

	l := NewLoginer(11)
	l.Login()
}

type Stringer interface {
	String() string
}

type person struct {
	name string
	age  int
}

func (p *person) String() string {
	return p.name + "'s age is " + strconv.Itoa(p.age)
}

func modify(p person) {
	p.age = p.age + 10
}

func modify2(p *person) {
	p.age = p.age + 10
}

func print(a ...interface{}) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func NewLoginer(i int) Loginer {
	return defaultLogin(i)
}

type Loginer interface {
	Login()
}

type defaultLogin int

func (d defaultLogin) Login() {
	fmt.Println("login in...", d)
}
