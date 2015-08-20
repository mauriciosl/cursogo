package main

import "github.com/mauriciosl/cursogo/server"

func main() {
	server := server.NewCourseServer()
	server.Run()
}
