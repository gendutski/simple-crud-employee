package repository

import (
	"fmt"
	"math"
	"simple-crud-employee/internal/entity"

	"gorm.io/gorm"
)

func InitEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

type EmployeeRepository struct {
	db *gorm.DB
}

// insert data into employee table
func (r *EmployeeRepository) Create(payload *entity.Employee) error {
	return r.db.Save(payload).Error
}

// update existing employee data
func (r *EmployeeRepository) Update(payload entity.Employee, data *entity.Employee) error {
	return r.db.Model(data).Updates(payload).Error
}

// delete existing data
func (r *EmployeeRepository) Delete(employee *entity.Employee) error {
	return r.db.Delete(employee).Error
}

// get query list of employees
func (r *EmployeeRepository) GetList(req *entity.QueryRequest) (*entity.EmployeeListResponse, error) {
	var employees []*entity.Employee
	var total int64

	qry := r.db.Model(&entity.Employee{})
	if req.EmployeeID != "" {
		qry = qry.Where("employee_id = ?", req.EmployeeID)
	}
	if req.FullName != "" {
		qry = qry.Where("full_name like ?", fmt.Sprintf("%%%s%%", req.FullName))
	}
	if req.Address != "" {
		qry = qry.Where("address like ?", fmt.Sprintf("%%%s%%", req.Address))
	}

	// get total records
	err := qry.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// validate pagination
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize < 10 {
		req.PageSize = 10
	} else if req.PageSize > 100 {
		req.PageSize = 100
	}

	// get total pages
	pages := math.Ceil(float64(total) / float64(req.PageSize))

	// get offset
	offset := (req.Page - 1) * req.PageSize

	// get employee list
	err = qry.Offset(offset).Limit(req.PageSize).Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return &entity.EmployeeListResponse{
		Employees: employees,
		Pages:     int(pages),
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}

// get single employee by employeeID
func (r *EmployeeRepository) GetDetail(employeeID string) (*entity.Employee, error) {
	var employee entity.Employee
	err := r.db.Where("employee_id = ?", employeeID).First(&employee).Error
	return &employee, err
}
