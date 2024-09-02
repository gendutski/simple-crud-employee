package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// setup routes endpoint
func SetupRoutes(server *Server) {
	// index page
	server.Router.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	// create employee data
	server.Router.POST("/create", server.EmployeeHandler.Create)

	// update eployee data
	server.Router.PUT("/update/:employeeID", server.EmployeeHandler.Update)

	// delete employee data
	server.Router.DELETE("/delete/:employeeID", server.EmployeeHandler.Delete)

	// get employee list
	server.Router.GET("/list", server.EmployeeHandler.Get)
}
