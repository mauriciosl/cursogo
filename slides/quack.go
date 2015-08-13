package main

import "fmt"

// interface START OMIT
type Quacker interface {
	Quack() string
}

// interface END OMIT

// implement START OMIT
type Pato struct{}

func (p Pato) Quack() string {
	return "Quack!"
}

type Marreco struct{}

func (m Marreco) Quack() string {
	return "Quein!"
}

// implement END OMIT

// use START OMIT
func Noise(q Quacker) {
	fmt.Println(q.Quack())
}

func main() {
	p := Pato{}
	m := Marreco{}
	Noise(p)
	Noise(m)
}

// use END OMIT
