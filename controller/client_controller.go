package controller

import (
	"net/http"

	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

type ClientRequest struct {
	ClientName   string  `json:"clientName" binding:"required,min=3"`
	ClientRegion string  `json:"clientRegion" binding:"required,min=3"`
	GST          float32 `json:"GST" binding:"required"`
}

// display client table
func GetClients(c *gin.Context) {
	var clients []models.Client
	models.DB.Find(&clients)
	c.JSON(http.StatusOK, gin.H{"data": clients})
}

func CreateClient(c *gin.Context) {
	var input ClientRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := models.Client{
		ClientName:  string(input.ClientName),
		ClientRegin: string(input.ClientRegion),
		GST:         float32(input.GST),
	}
	models.DB.Create(&client)
	c.JSON(http.StatusOK, gin.H{"data": client})
}

//find a client
func FindClientById(c *gin.Context) {
	var client models.Client
	err := models.DB.Where("client_id = ?", c.Param("clientId")).First(&client).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": client})
}

func UpdateClient(c *gin.Context) {
	var client models.Client
	if err := models.DB.Where("client_id = ?", c.Param("clientId")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var input ClientRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&client).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": client})
}

//delete client
func DeleteClient(c *gin.Context) {
	var client models.Client
	if err := models.DB.Where("client_id = ?", c.Param("clientId")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&client)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
