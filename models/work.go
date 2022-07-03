package models

type Work struct {
	Year       int      `json:"year"`
	Month      int      `json:"month"`
	Days       int      `json:"days"`
	EmployeeID uint     `json:"employeeId"`
	ClientID   uint     `json:"clientId"`
	Employee   Employee `json:"-"`
	Client     Client   `json:"-"`
}
