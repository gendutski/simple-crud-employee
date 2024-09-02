package entity

type QueryRequest struct {
	EmployeeID string `json:"employeeID" query:"employeeID"`
	FullName   string `json:"fullName" query:"fullName"`
	Address    string `json:"address" query:"address"`
	Page       int    `json:"page" query:"page"`
	PageSize   int    `json:"pageSize" query:"pageSize"`
}
