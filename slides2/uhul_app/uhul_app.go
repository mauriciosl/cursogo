package main

import (
	"fmt"
	"net/http"
)

// struct START OMIT
type UhulGreeter struct {
	Message string
}

func (u UhulGreeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s!", u.Message)
}

// struct END OMIT

// main START OMIT
func main() {
	app := UhulGreeter{Message: "Viajar"}
	http.Handle("/", app)
	http.ListenAndServe(":6680", nil)
}

// main END OMIT
