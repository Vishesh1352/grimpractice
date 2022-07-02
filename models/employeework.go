package models

import "time"

type EmployeeWork struct {
	EmployeeID uint      `json:"employeeID" `
	ClientID   uint      `json:"clientID"`
	WorkDate   time.Time `json:"workDate" binding:"required" time_format:"2006-01-02"`
	WorkHours  float32   `json:"workHours"`
}
