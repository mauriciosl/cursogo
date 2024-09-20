package main

import (
	"fmt"
	"net/http"
)

// struct START OMIT
type AuuGreeter struct {
	Message string
}

func (u AuuGreeter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s!", u.Message)
}

// struct END OMIT

// main START OMIT
func main() {
	app := AuuGreeter{Message: "Viajar"}
	http.Handle("/", app)
	http.ListenAndServe(":6680", nil)
}

// main END OMIT
