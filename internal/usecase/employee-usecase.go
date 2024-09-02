package usecase

import (
	"errors"
	"net/http"
	"simple-crud-employee/internal/entity"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	mysqlDuplicateErrorNum uint16 = 1062
)

type EmployeeRepo interface {
	// insert data into employee table
	Create(payload *entity.Employee) error
	// update existing employee data
	Update(payload entity.Employee, data *entity.Employee) error
	// delete existing data
	Delete(employee *entity.Employee) error
	// get query list of employees
	GetList(req *entity.QueryRequest) (*entity.EmployeeListResponse, error)
	// get single employee by employeeID
	GetDetail(employeeID string) (*entity.Employee, error)
}

func InitEmployeeUsecase(repo EmployeeRepo) *EmployeeUsecase {
	return &EmployeeUsecase{
		repo:     repo,
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

type EmployeeUsecase struct {
	repo     EmployeeRepo
	validate *validator.Validate
}

func (uc *EmployeeUsecase) Create(payload *entity.Employee) error {
	// validate payload
	err := uc.validate.Struct(payload)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err,
		}
	}

	// create employee
	err = uc.repo.Create(payload)

	if err != nil {
		// duplicate create, return 409 conflict
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr != nil && mysqlErr.Number == mysqlDuplicateErrorNum {
			return &echo.HTTPError{
				Code:    http.StatusConflict,
				Message: err,
			}
		}
		// internal server error
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}
	return nil
}

func (uc *EmployeeUsecase) Update(employeeID string, payload entity.Employee) error {
	// get employee detail
	employee, err := uc.repo.GetDetail(employeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// employee not found
			return &echo.HTTPError{
				Code:    http.StatusNotFound,
				Message: err,
			}
		}
		// internal server error
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}
	err = uc.repo.Update(payload, employee)
	if err != nil {
		// duplicate create, return 409 conflict
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr != nil && mysqlErr.Number == mysqlDuplicateErrorNum {
			return &echo.HTTPError{
				Code:    http.StatusConflict,
				Message: err,
			}
		}
		// internal server error
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}
	return nil
}

func (uc *EmployeeUsecase) Delete(employeeID string) error {
	// get employee detail
	employee, err := uc.repo.GetDetail(employeeID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// employee not found
			return nil
		}
		// internal server error
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err,
		}
	}
	return uc.repo.Delete(employee)
}

func (uc *EmployeeUsecase) Get(req *entity.QueryRequest) (*entity.EmployeeListResponse, error) {
	return uc.repo.GetList(req)
}
