package main

import (
	"fmt"
	"net/http"
)

// func START OMIT
func AuuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Auu!")
}

// func END OMIT

// main START OMIT
func main() {
	http.HandleFunc("/", AuuHandler)
	http.ListenAndServe(":6680", nil)
}

// main END OMIT
