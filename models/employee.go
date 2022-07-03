package models

type Employee struct {
	EmployeeID   uint    `json:"employeeId" gorm:"primary_key;auto_increment;not_null"`
	EmployeeName string  `json:"employeeName"`
	EmployeeRate float32 `json:"employeeRate"`
	ClientID     uint    `json:"clientId"`
	Client       Client  `json:"-"`
}
