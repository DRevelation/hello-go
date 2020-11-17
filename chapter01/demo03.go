package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s1 = "abc"

	// 浮点型默认float64
	var rate64 = 3.14
	var rate32 float32 = 3.14

	var ptr = new(int)

	fmt.Println(reflect.TypeOf(rate64))
	fmt.Println(reflect.TypeOf(rate32))
	fmt.Println("ptr type=", reflect.TypeOf(ptr))
	fmt.Println("ptr address=", ptr)
	fmt.Println("ptr value=", *ptr)

	fmt.Println(s1)
	fmt.Println("s1 address=", &s1)
	fmt.Println("s1 value=", *(&s1))

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187

	fmt.Println("myfloat:", myfloat01)
	fmt.Println("myfloat:", myfloat01+5)
	// float精度问题，结果是false
	fmt.Println(myfloat02 == myfloat01+5)

	//5+1+(3*2)
	var country = "hello,中国"
	fmt.Println(len(country))

	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6:8]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))

	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}
