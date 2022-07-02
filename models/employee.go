package models

type Employee struct {
	EmployeeID   uint   `json:"Employeeid" gorm:"primary_key;auto_increment;not_null"`
	EmployeeName string `json:"EmployeeName"`
	EmployeeRate int    `json:"EmployeeRate"`
}
