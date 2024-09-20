package main

import "fmt"

// START OMIT
func mudaMap(p map[string]int) {
	p["ei"] = 2
}

func main() {
	x := map[string]int{"score": 10}
	z := x
	z["score"] = 7
	fmt.Println(x, z) // ???
	z["zz"] = 0
	fmt.Println(x, z) // ???
	mudaMap(z)
	fmt.Println(x, z) // ???
}

// END OMIT
