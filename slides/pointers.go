package main

import "fmt"

type Point struct {
	X, Y int
}

// START OMIT
func (p *Point) Mul(v int) {
	p.X, p.Y = p.X*v, p.Y*v
}

func (p Point) CopyMul(v int) Point {
	p.X, p.Y = p.X*v, p.Y*v
	return p
}

func main() {
	p := &Point{10, 5}
	p.Mul(2)
	fmt.Println(p)
	fmt.Println(p.CopyMul(2))
	fmt.Println(p)
}

// END OMIT
