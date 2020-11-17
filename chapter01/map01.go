package main

import (
	"fmt"
	"sort"
)

func main() {
	var map01 = make(map[string]int)

	var names []string

	map01["d"] = 4
	map01["a"] = 1
	map01["b"] = 2
	map01["c"] = 3

	for k, v := range map01 {
		names = append(names, k)
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	fmt.Println("=======================")
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("key=%s, value=%d\n", name, map01[name])
	}

	type person struct {
		nam
	}
}
