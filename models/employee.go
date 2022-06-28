package models

type employeee struct {
	employeeID   uint   `json : "employeeID" gorm :"primary_key"`
	employeeName string `json:"employeeName"`
	employeeRate int    `json:"employeeRate"`
}
