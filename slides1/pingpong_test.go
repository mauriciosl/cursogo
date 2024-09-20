package main

import "testing"

func TestPingPong(t *testing.T) {

	if got := PingPong(3); got != "ping" {
		t.Errorf("PingPong(3) era pra ser ping, veio %s", got)
	}
	if got := PingPong(5); got != "pong" {
		t.Errorf("PingPong(5) era pra ser pong, veio %s", got)
	}
	if got := PingPong(15); got != "ping pong" {
		t.Errorf("PingPong(15) era pra ser ping pong, veio %s", got)
	}
	if got := PingPong(1); got != "1" {
		t.Errorf("PingPong(1) era pra ser 1, veio %s", got)
	}
}
