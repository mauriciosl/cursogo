package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (c *CourseServer) setupRoutes() {
	c.router = mux.NewRouter()
	c.router.HandleFunc("/healthcheck", c.healthcheck)
}

func (c *CourseServer) healthcheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
