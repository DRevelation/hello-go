package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [5]int
	fmt.Println("数组初始化", arr)
	fmt.Println(reflect.TypeOf(arr))

	var sli []int
	fmt.Println("切片初始化", sli)
	fmt.Println(reflect.TypeOf(sli))

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	var sli2 = make([]int, 5, 10)
	fmt.Println(sli2)
}
