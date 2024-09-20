package main

import (
	"fmt"
	"net/http"
)

func AuuHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Auu!")
}

func main() {
	http.HandleFunc("/", AuuHandler)
	http.ListenAndServe(":6680", nil)
}
