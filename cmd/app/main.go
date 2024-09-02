package main

import "simple-crud-employee/internal/infrastructure/server"

func main() {
	// set server
	srv := server.InitServer()

	// set router
	server.SetupRoutes(srv)

	// set templates
	server.InitTemplates(srv)

	// run http
	srv.Run()
}
