package controller

import (
	"net/http"
	"strconv"
	"time"

	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

func GetBillForAnEmployee(c *gin.Context) {

	year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
	month := c.DefaultQuery("month", "")

	if month == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MONTH_INFO_REQUIRED"})
		return
	}

	employeeId := c.Param("employeeId")
	// eId, err := strconv.ParseUint(c.Param("employeeId"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_INPUT"})
	// 	return
	// }

	var WorkModel []models.Work

	if err := models.DB.Where("employee_id = ? and year = ? and month = ?", employeeId, year, month).First(&WorkModel, employeeId, year, month).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_ERROR"})
		return
	}

	clientId := WorkModel[0].ClientID
	workingDays := WorkModel[0].Days

	var ClientModel []models.Client

	if err := models.DB.Where("client_id = ?", clientId).First(&ClientModel, clientId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_ERROR"})
		return
	}

	var EmployeeModel []models.Employee

	if err := models.DB.Where("employee_id = ?", employeeId).First(&EmployeeModel, employeeId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL_ERROR"})
		return
	}

	employeeRate := EmployeeModel[0].EmployeeRate

	gst := ClientModel[0].GST

	billAmount := employeeRate * float32(workingDays)

	gstAmount := billAmount * gst / 100

	totalAmount := billAmount + gstAmount

	c.JSON(http.StatusOK, gin.H{"totalBill": totalAmount})
}

func GetBillForAClient(c *gin.Context) {

	// year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
	// month := c.DefaultQuery("month", "")

	// if month == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "MONTH_INFO_REQUIRED"})
	// 	return
	// }

	// cId, err := strconv.ParseUint(c.Param("clientId"), 10, 32)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_INPUT"})
	// 	return
	// }

	// var clients []models.Client
	// models.DB.Find(&clients)
	c.JSON(http.StatusOK, gin.H{"data": "clients"})
}
