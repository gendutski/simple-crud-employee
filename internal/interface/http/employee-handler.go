package http

import (
	"net/http"
	"simple-crud-employee/internal/entity"

	"github.com/labstack/echo/v4"
)

type EmployeeUsecase interface {
	Create(payload *entity.Employee) error
	Update(employeeID string, payload entity.Employee) error
	Delete(employeeID string) error
	Get(req *entity.QueryRequest) (*entity.EmployeeListResponse, error)
}

func InitEmployeeHandler(employeeUsecase EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{employeeUsecase}
}

type EmployeeHandler struct {
	usecase EmployeeUsecase
}

func (h *EmployeeHandler) Create(e echo.Context) error {
	// bind payload
	payload := new(entity.Employee)
	err := e.Bind(payload)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	// create employee
	err = h.usecase.Create(payload)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, payload)
}

func (h *EmployeeHandler) Update(e echo.Context) error {
	// get param
	employeID := e.Param("employeeID")
	if employeID == "" {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "employeID must not empty",
		}
	}

	// bind payload
	var payload entity.Employee
	err := e.Bind(&payload)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	// update employee
	err = h.usecase.Update(employeID, payload)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, payload)
}

func (h *EmployeeHandler) Delete(e echo.Context) error {
	// get param
	employeID := e.Param("employeeID")
	if employeID == "" {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "employeID must not empty",
		}
	}

	// delete employee
	err := h.usecase.Delete(employeID)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusNoContent, nil)
}

func (h *EmployeeHandler) Get(e echo.Context) error {
	// bind request
	req := new(entity.QueryRequest)
	err := e.Bind(req)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}

	// get list data
	resp, err := h.usecase.Get(req)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, resp)
}
