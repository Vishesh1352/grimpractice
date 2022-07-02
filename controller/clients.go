package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/grimpractice/models"

)

// display client table
func FindClients(c *gin.Context) {
	var clients []models.Client

	models.DB.Find(&clients)

	c.JSON(http.StatusOK, gin.H{"data": clients})
}

//create client
type CreateClient struct {
	ClientName   string `json:"clientName" binding:"required"`
	ClientRegion string `json:"clientRegion" binding:"required"`
	GST          int    `json:"GST" binding:"required"`
}

func CreateClients(c *gin.Context) {

	var input CreateClient
	//check what shoudbingjson do????
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := models.Client{
		ClientName:  string(input.ClientName),
		ClientRegin: string(input.ClientRegion),
		GST:         uint(input.GST),
	}
	models.DB.Create(&client)
	c.JSON(http.StatusOK, gin.H{"data": client})
}

//find a client
func FindClientsid(c *gin.Context) {

	var client models.Client

	err := models.DB.Where("client_id = ?", c.Param("id")).First(&client).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": client})
}

//Update Client
type UpdateClientInput struct {
	ClientName  string `json:"ClientName"`
	ClientRegin string `json:"ClientRegion"`
	GST         int    `json:"GST"`
}

func UpdateClient(c *gin.Context) {

	var client models.Client
	if err := models.DB.Where("client_id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateClientInput
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
	if err := models.DB.Where("client_id = ?", c.Param("id")).First(&client).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&client)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
