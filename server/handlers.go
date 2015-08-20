package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mauriciosl/cursogo/keydb"
)

func (c *CourseServer) setupRoutes() {
	c.router = mux.NewRouter()
	c.router.HandleFunc("/healthcheck", c.healthcheck)
	c.router.Handle("/db/{key}", keydb.NewDB())
}

func (c *CourseServer) healthcheck(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
