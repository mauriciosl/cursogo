package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CourseServer struct {
	router *mux.Router
}

func NewCourseServer() *CourseServer {
	c := &CourseServer{}
	c.setupRoutes()
	return c
}
func (c *CourseServer) Run() {
	http.ListenAndServe(":6680", c.router)
}
