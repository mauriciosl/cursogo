package main

import "fmt"

// START OMIT
func main() {
	x := [3]string{"pera", "coco", "uva"}
	y := x[:]         // Cria um slice a partir da array
	fmt.Println(x, y) // [pera coco uva] [pera coco uva]
	y[0] = "caju"
	fmt.Println(x, y) // [caju coco uva] [caju coco uva]
	y = append(y, "kiwi")
	fmt.Println(x, y) // [caju coco uva] [caju coco uva kiwi]
	y[0] = "jaca"
	fmt.Println(x, y) // [caju coco uva] [jaca coco uva kiwi]
}

// END OMIT
