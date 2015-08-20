package main

import "fmt"

// type START OMIT
type Mes int

// type END OMIT

const (
	Jan Mes = iota
	Fev
	Abr
	Mai
	Jun
	Jul
	Ago
	Out
	Set
	Nov
	Dez
)

var nomes = [...]string{"Jan", "Fev", "Abr", "Mai", "Jun", "Jul", "Ago", "Out", "Set", "Nov", "Dez"}

// String START OMIT

func (m Mes) String() string {
	// String END OMIT
	return nomes[m]
}

// main START OMIT
func main() {
	for m := Jan; m <= Dez; m++ {
		fmt.Println(m.String())
	}
}

// main END OMIT
