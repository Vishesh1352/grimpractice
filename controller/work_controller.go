package controller

import (
	"net/http"
	"strconv"

	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

type WorkRequest struct {
	Year  int `json:"year"  binding:"required,max=4"`
	Month int `json:"month" binding:"required,max=2"`
	Days  int `json:"days" binding:"required,max=2"`
}

func AddWork(c *gin.Context) {
	var workRequest WorkRequest
	if err := c.ShouldBindJSON(&workRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cId, err := strconv.ParseUint(c.Param("clientId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_INPUT"})
		return
	}

	eId, err := strconv.ParseUint(c.Param("employeeId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_INPUT"})
		return
	}

	Work := models.Work{
		EmployeeID: uint(eId),
		ClientID:   uint(cId),
		Year:       workRequest.Year,
		Month:      workRequest.Month,
		Days:       workRequest.Days,
	}
	models.DB.Create(&Work)
	c.JSON(http.StatusOK, gin.H{"data": Work})
}

func UpdateWork(c *gin.Context) {

	var WorkModel models.Work
	employe_id := c.Param("id")
	client_id := c.Param("cid")
	date := c.Param("date")

	if err := models.DB.Where("employee_id = ?", employe_id,
		"client_id = ?", client_id).First(&WorkModel, employe_id, client_id, date).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_DATA"})
		return
	}

	var input WorkRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&WorkModel).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": input})
}
