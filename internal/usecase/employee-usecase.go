package usecase

import "simple-crud-employee/internal/entity"

type EmployeeRepo interface {
	// insert data into employee table
	Create(payload *entity.Employee) error
	// update existing employee data
	Update(payload entity.Employee, data *entity.Employee) error
	// delete existing data
	Delete(employeeID string) error
	// get query list of employees
	GetList(employeeID, fullName, address string) (*entity.EmployeeListResponse, error)
	// get single employee by employeeID
	GetDetail(employeeID string) (*entity.Employee, error)
}

func InitEmployeeUsecase(repo EmployeeRepo) *EmployeeUsecase {
	return &EmployeeUsecase{repo}
}

type EmployeeUsecase struct {
	repo EmployeeRepo
}

func (uc *EmployeeUsecase) Create(payload *entity.Employee) error {

}

func (uc *EmployeeUsecase) Update(employeeID string, payload *entity.Employee) error {

}

func (uc *EmployeeUsecase) Delete(employeeID string) error {

}

func (uc *EmployeeUsecase) Get(employeeID, fullName, address string) ([]*entity.Employee, error) {

}
