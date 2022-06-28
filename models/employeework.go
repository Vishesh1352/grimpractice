package models

type employeeWork struct {
	employeeName string  `json :"employeeName"`
	employeeID   uint    `json:"employeeID" `
	clientID     uint    `json :"clientID"`
	clientName   string  `json:"clientName"`
	workDate     string  `json:"workDate"`
	workHours    float32 `json:"workHours"`
}
