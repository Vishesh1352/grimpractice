package controller

import (
	"net/http"

	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

type EmployeeRequest struct {
	EmployeeName string  `binding:"required"`
	EmployeeRate float32 `binding:"required"`
	ClientID     int     `binding:"required"`
}

// display Employee table
func GetEmployees(c *gin.Context) {
	var Employees []models.Employee
	models.DB.Find(&Employees)
	c.JSON(http.StatusOK, gin.H{"data": Employees})
}

func CreateEmployee(c *gin.Context) {
	var input EmployeeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Employee := models.Employee{
		EmployeeName: string(input.EmployeeName),
		EmployeeRate: float32(input.EmployeeRate),
		ClientID:     uint(input.ClientID),
	}
	models.DB.Create(&Employee)
	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

//find a Employee
func FindEmployeeById(c *gin.Context) {
	var Employee models.Employee
	err := models.DB.Where("employee_id = ?", c.Param("employeeId")).First(&Employee).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

func UpdateEmployee(c *gin.Context) {
	var Employee models.Employee
	if err := models.DB.Where("employee_id = ?", c.Param("employeeId")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input EmployeeRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&Employee).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

//delete Employee
func DeleteEmployee(c *gin.Context) {
	var Employee models.Employee
	if err := models.DB.Where("employee_id = ?", c.Param("employeeId")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&Employee)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
