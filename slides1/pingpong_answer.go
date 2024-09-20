package main

import (
	"fmt"
	"strconv"
)

func PingPong(x int) string {
	if x%3+x%5 == 0 {
		return "ping pong"
	} else if x%3 == 0 {
		return "ping"
	} else if x%5 == 0 {
		return "pong"
	}
	return strconv.Itoa(x)
}

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("%d: %s\n", i, PingPong(i))
	}
}
