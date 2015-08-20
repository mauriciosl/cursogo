package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func say(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	time.Sleep(time.Duration(i) * time.Second)
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
