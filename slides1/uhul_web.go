package main

import (
	"fmt"
	"net/http"
)

func UhulHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Uhul!")
}

func main() {
	http.HandleFunc("/", UhulHandler)
	http.ListenAndServe(":6680", nil)
}
