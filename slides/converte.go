package main

import "fmt"

func main() {
	var x interface{}
	x = 2
	// y := x + 2
	z := x.(int)
	fmt.Println(z + 2)
	v, ok := x.(string)
	fmt.Println(v, ok)
}
