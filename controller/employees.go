package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/grimpractice/models"

)

// display Employee table
func FindEmployees(c *gin.Context) {
	
	var Employees []models.Employee

	models.DB.Find(&Employees)

	c.JSON(http.StatusOK, gin.H{"data": Employees})
}

//create Employee
type CreateEmployee struct {
	EmployeeName string `json:"EmployeeName" binding:"required"`
	EmployeeRate int    `json:"EmployeeRate" binding:"required"`
}

func CreateEmployees(c *gin.Context) {

	var input CreateEmployee
	//check what shoudbingjson do????
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Employee := models.Employee{
		EmployeeName: string(input.EmployeeName),
		EmployeeRate: int(input.EmployeeRate),
	}
	models.DB.Create(&Employee)
	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

//find a Employee
func FindEmployeesid(c *gin.Context) {

	var Employee models.Employee

	err := models.DB.Where("employee_id = ?", c.Param("id")).First(&Employee).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employee})
}

//Update Employee
type UpdateEmployeeInput struct {
	EmployeeName  string `json:"EmployeeName"`
	EmployeeRegin string `json:"EmployeeRegion"`
	GST           int    `json:"GST"`
}

func UpdateEmployee(c *gin.Context) {

	var Employee models.Employee
	if err := models.DB.Where("employee_id = ?", c.Param("id")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateEmployeeInput
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
	if err := models.DB.Where("employee_id = ?", c.Param("id")).First(&Employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Employee)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

