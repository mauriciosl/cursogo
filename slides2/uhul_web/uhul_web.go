package main

import (
	"fmt"
	"net/http"
)

// func START OMIT
func UhulHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Uhul!")
}

// func END OMIT

// main START OMIT
func main() {
	http.HandleFunc("/", UhulHandler)
	http.ListenAndServe(":6680", nil)
}

// main END OMIT
