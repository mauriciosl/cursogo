package main

import "fmt"

func say(s string) {
	fmt.Print(s)
}

func main() {
	defer say(" possivel\n")
	say("Viajar é")
	// panic("socorro!")
}
