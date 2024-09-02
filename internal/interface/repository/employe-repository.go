package repository

import (
	"simple-crud-employee/internal/entity"

	"gorm.io/gorm"
)

func InitEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

type EmployeeRepository struct {
	db *gorm.DB
}

func (r *EmployeeRepository) Create(payload *entity.Employee) error {

}

func (r *EmployeeRepository) Update(employeeID string, payload *entity.Employee) error {

}

func (r *EmployeeRepository) Delete(employeeID string) error {

}

func (r *EmployeeRepository) Get(employeeID, fullName, address string) ([]*entity.Employee, error) {

}
