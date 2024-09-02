package entity

type EmployeeListResponse struct {
	Employees []*Employee
	Pages     int
	Page      int
	Limit     int
}
