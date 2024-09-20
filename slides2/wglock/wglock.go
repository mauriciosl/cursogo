package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func say(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	// Espera entre 0.5 e 2.5 segundos
	time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
	fmt.Printf("done %d\n", i)
}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go say(wg, i)
	}
	wg.Wait()
}

// END OMIT
