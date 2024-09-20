package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(s)
	}
}

func main() {
	go say(" Stone\n")
	say("Conta")
	time.Sleep(200 * time.Millisecond)
}
