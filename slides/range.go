package main

import "fmt"

// START OMIT
func main() {
	fib := []int{1, 2, 3, 5, 8, 13, 21}
	for index, value := range fib {
		fmt.Println(index, "=>", value)
	}
	id2name := map[int]string{
		1: "Blue Tree",
		2: "Lua de Tomate",
	}
	fmt.Println(id2name)
	name2id := make(map[string]int)
	for key, value := range id2name {
		name2id[value] = key
	}
	fmt.Println(name2id)
}

// END OMIT
