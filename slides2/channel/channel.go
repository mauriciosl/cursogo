package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func worker(output chan int) {
	i := rand.Intn(5)
	fmt.Printf("start %d\n", i)
	time.Sleep(time.Duration(i) * time.Second)
	output <- i * 2
}

func main() {
	outputs := make(chan int)
	for i := 0; i < 4; i++ {
		go worker(outputs)
	}
	for i := 0; i < 4; i++ {
		fmt.Println(<-outputs)
	}
}

// END OMIT
