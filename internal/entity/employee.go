package entity

type Employee struct {
	EmployeeID string `gorm:"primaryKey;size:4" json:"employeeID" validate:"required,max=4"`
	FullName   string `gorm:"size:100;not null" json:"fullName" validate:"required,max=100"`
	Address    string `gorm:"type:text;not null" json:"address" validate:"required"`
}
