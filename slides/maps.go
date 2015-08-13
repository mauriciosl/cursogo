package main

import "fmt"

// START OMIT
func blah(p map[string]int) {
	p["ei"] = 2
}

func main() {
	x := map[string]int{"hu": 10}
	z := x
	z["hu"] = 11
	fmt.Println(x, z) // ???
	z["zz"] = 0
	fmt.Println(x, z) // ???
	blah(z)
	fmt.Println(x, z) // ???
}

// END OMIT
