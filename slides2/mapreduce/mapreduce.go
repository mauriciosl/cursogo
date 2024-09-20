package main

import (
	"fmt"
	"math/rand"
	"time"
)

// worker START OMIT
func worker(i int, input chan int, output chan int) {
	for n := range input {
		time.Sleep(time.Duration(rand.Intn(1000)+1) * time.Millisecond)
		fmt.Printf("worker[%d]: processing %d \n", i, n)
		output <- n
	}
	fmt.Printf("worker[%d]: exiting\n", i)
}

// worker END OMIT

// main START OMIT
func main() {
	N := 20
	input := make(chan int, N)
	output := make(chan int)
	for i := 0; i < 4; i++ {
		go worker(i, input, output)
	}
	for i := 0; i < N; i++ {
		input <- i
	}
	close(input)
	for i := 0; i < N; i++ {
		fmt.Println(<-output)
	}
}

// main END OMIT
