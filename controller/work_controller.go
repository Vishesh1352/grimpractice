package controller

import (
	"net/http"
	"strconv"

	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

type WorkRequest struct {
	Year  int `json:"year"  binding:"required"`
	Month int `json:"month" binding:"required"`
	Days  int `json:"days" binding:"required"`
}

func AddWork(c *gin.Context) {
	var workRequest WorkRequest
	//check what shoudbingjson do????
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

// // display EmployeeWork table
// func FindEmployeeWorks(c *gin.Context) {
// 	var EmployeeWork []models.EmployeeWork
// 	models.DB.Find(&EmployeeWork)
// 	c.JSON(http.StatusOK, gin.H{"data": EmployeeWork})
// }

// func CreateEmployeeWork(c *gin.Context) {
// 	var input Work
// 	//check what shoudbingjson do????
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	Work := models.EmployeeWork{
// 		EmployeeID: uint(input.EmployeeID),
// 		ClientID:   uint(input.ClientID),
// 		WorkDate:   time.Time(input.WorkDate),
// 		WorkHours:  float32(input.WorkHours),
// 	}
// 	models.DB.Create(&Work)
// 	c.JSON(http.StatusOK, gin.H{"data": Work})
// }

// //find a EmployeeWork
// func FindEmployeeWorkid(c *gin.Context) {

// 	var EmployeeWork models.EmployeeWork

// 	err := models.DB.Where("employee_id = ?", c.Param("id")).First(&EmployeeWork).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": EmployeeWork})
// }

// //find a ClientWork
// func FindClientWorkid(c *gin.Context) {

// 	var ClientWork models.EmployeeWork

// 	err := models.DB.Where("client_id = ?", c.Param("id")).First(&ClientWork).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": ClientWork})
// }

// //Update EmployeeWork
// type UpdateEmployeeWorkInput struct {
// 	EmployeeID uint      `json:"employeeID" `
// 	ClientID   uint      `json:"clientID"`
// 	WorkDate   time.Time `json:"workDate"  binding:"required" time_format:"2006-01-02"`
// 	WorkHours  float32   `json:"workHours" binding:"required"`
// }

// func UpdateEmployeeWork(c *gin.Context) {

// 	var EmployeeWork models.EmployeeWork
// 	employe_id := c.Param("id")
// 	client_id := c.Param("cid")
// 	date := c.Param("date")

// 	if err := models.DB.Where("employee_id = ?", employe_id,
// 		"client_id = ?", client_id,
// 		"date = ?", date).First(&EmployeeWork, employe_id, client_id, date).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	var input UpdateEmployeeWorkInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Model(&EmployeeWork).Updates(input)

// 	c.JSON(http.StatusOK, gin.H{"data": EmployeeWork})
// }

//delete EmployeeWork
// func DeleteEmployeeWork(c *gin.Context) {

// 	var EmployeeWork models.EmployeeWork
// 	if err := models.DB.Where("employee_id = ?", c.Param("id")).First(&EmployeeWork).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	models.DB.Delete(&EmployeeWork)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }

// //delete ClientWork
// func DeleteClientWork(c *gin.Context) {

// 	var ClientWork models.EmployeeWork
// 	if err := models.DB.Where("client_id = ?", c.Param("id")).First(&ClientWork).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	models.DB.Delete(&ClientWork)

// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }

//delete specific Entity
func DeleteWork(c *gin.Context) {

	var WorkModel models.Work
	client_id := c.Param("cid")
	date := c.Param("date")
	employee_id := c.Param("id")
	if err := models.DB.Where("employee_id = ?", employee_id,
		"client_id = ?", client_id).First(&WorkModel, employee_id, client_id, date).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_DATA"})
		return
	}

	models.DB.Delete(&WorkModel)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
