package main

import "fmt"

func say(s string) {
	fmt.Print(s)
}

func main() {
	defer say(" Stone\n")
	say("Conta")
	// panic("socorro!")
}
