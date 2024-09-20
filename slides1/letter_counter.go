package main

import "fmt"

// START OMIT
type LetterCountMap map[rune]int

func (m LetterCountMap) Add(c rune) {
	actual, exists := m[c]
	if !exists {
		m[c] = 1
		return
	}
	m[c] = actual + 1
}

func main() {
	counter := LetterCountMap{}
	for _, y := range "concorrência" {
		counter.Add(y)
	}
	for letter, count := range counter {
		fmt.Printf("%q: %d\n", letter, count)
	}
}

// END OMIT
