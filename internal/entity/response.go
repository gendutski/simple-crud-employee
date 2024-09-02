package entity

type EmployeeListResponse struct {
	Employees []*Employee `json:"data"`
	Pages     int         `json:"pages"`
	Page      int         `json:"page"`
	PageSize  int         `json:"pageSize"`
}
