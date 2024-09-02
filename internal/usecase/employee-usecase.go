package usecase

import "simple-crud-employee/internal/entity"

type EmployeeRepo interface {
	Create(payload *entity.Employee) error
	Update(employeeID string, payload *entity.Employee) error
	Delete(employeeID string) error
	Get(employeeID, fullName, address string) ([]*entity.Employee, error)
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
