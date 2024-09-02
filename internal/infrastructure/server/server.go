package server

import (
	"fmt"
	"simple-crud-employee/configs"
	handler "simple-crud-employee/internal/interface/http"
	"simple-crud-employee/internal/interface/repository"
	"simple-crud-employee/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Router          *echo.Echo
	EmployeeHandler *handler.EmployeeHandler
	config          configs.Server
}

// run http server
func (s *Server) Run() {
	s.Router.Logger.Fatal(s.Router.Start(fmt.Sprintf(":%d", s.config.Port)))
}

func InitServer() *Server {
	// get server config
	cfg := configs.InitServerConfig()

	// get database config
	db := configs.InitDatabaseConfig().ConnectDB()

	// set repo
	employeeRepo := repository.InitEmployeeRepository(db)

	// set usecase
	employeeUsecase := usecase.InitEmployeeUsecase(employeeRepo)

	// set handler
	employeeHandler := handler.InitEmployeeHandler(employeeUsecase)

	// settup echo router
	e := echo.New()

	return &Server{
		Router:          e,
		EmployeeHandler: employeeHandler,
		config:          cfg,
	}
}
