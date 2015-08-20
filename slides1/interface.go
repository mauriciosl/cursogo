package main

import "fmt"

type Quacker interface {
	Quack() string
}

type Pato struct{}

func (p Pato) Quack() string {
	return "Quack!"
}

// START OMIT
func Reply(i interface{}) {
	switch i := i.(type) {
	default:
		fmt.Printf("Não sei tratar o tipo %T\n", i)
	case string:
		fmt.Printf("string: %s\n", i)
	case bool:
		fmt.Printf("bool: %t\n", i)
	case int:
		fmt.Printf("int: %d\n", i)
	case Quacker:
		fmt.Printf("%s\n", i.String())
	}
}

func main() {
	Reply("uhul")
	Reply(2)
	Reply(true)
	Reply(Pato{})
	Reply([]int{2})
}

// END OMIT
