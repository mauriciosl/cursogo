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
	go say(" possivel\n")
	say("Viajar é")
	time.Sleep(200 * time.Millisecond)
}
