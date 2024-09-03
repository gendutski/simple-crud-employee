package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-crud-employee/configs"
	"simple-crud-employee/internal/entity"
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

	// auto migrate database
	db.AutoMigrate(&entity.Employee{})

	// set repo
	employeeRepo := repository.InitEmployeeRepository(db)

	// set usecase
	employeeUsecase := usecase.InitEmployeeUsecase(employeeRepo)

	// set handler
	employeeHandler := handler.InitEmployeeHandler(employeeUsecase)

	// settup echo router
	e := echo.New()

	// custom http error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		errorHandler(e, err, c)
	}

	return &Server{
		Router:          e,
		EmployeeHandler: employeeHandler,
		config:          cfg,
	}
}

// echo error handler
func errorHandler(e *echo.Echo, err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	// error must type *echo.HTTPError
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	// set http status
	status := http.StatusText(he.Code)

	// check message interface
	message := he.Message
	switch m := he.Message.(type) {
	case string:
		message = echo.Map{"code": he.Code, "status": status, "message": m}
	case map[string][]*entity.ValidatorMessage:
		message = echo.Map{"code": he.Code, "status": status, "message": m}
	case json.Marshaler:
		// do nothing - this type knows how to format itself to JSON
	case error:
		message = echo.Map{"code": he.Code, "status": status, "message": m.Error()}
	}

	// Send response
	if c.Request().Method == http.MethodHead {
		err = c.NoContent(he.Code)
	} else {
		err = c.JSON(he.Code, message)
	}
	if err != nil {
		e.Logger.Error(err)
	}
}
