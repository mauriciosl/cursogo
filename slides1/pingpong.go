package main

import (
	"fmt"
	"strconv"
)

func PingPong(x int) string {
	return strconv.Itoa(x)
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%d: %s\n", i, PingPong(i))
	}
}
