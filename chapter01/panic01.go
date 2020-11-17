package main

import "fmt"

func main() {
	//panic("error")
	var i interface{} = 10
	var t1 = i.(int)
	fmt.Println(t1)

	fmt.Println("==============")
	t2 := i.(string)
	fmt.Println(t2)
}
