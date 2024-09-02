package http

import (
	"simple-crud-employee/internal/entity"

	"github.com/labstack/echo/v4"
)

type EmployeeUsecase interface {
	Create(payload *entity.Employee) error
	Update(employeeID string, payload entity.Employee) error
	Delete(employeeID string) error
	Get(employeeID, fullName, address string) ([]*entity.Employee, error)
}

func InitEmployeeHandler(employeeUsecase EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{employeeUsecase}
}

type EmployeeHandler struct {
	usecase EmployeeUsecase
}

func (h *EmployeeHandler) Create(e echo.Context) error {

}

func (h *EmployeeHandler) Update(e echo.Context) error {

}

func (h *EmployeeHandler) Delete(e echo.Context) error {

}

func (h *EmployeeHandler) Get(e echo.Context) error {

}
