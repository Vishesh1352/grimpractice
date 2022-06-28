package main

import (
	"github.com/gin-gonic/gin"

	"example.com/grimpractice/controller"
	"example.com/grimpractice/models"

)

func main() {
	r := gin.Default()

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	//})

	models.DBconnectDatabase()

	r.GET("/client",controller.FindClients)
	r.POST("/client",controller.CreateClients)
	r.GET("/client/:id",controller.FindClientsid)
	r.PATCH("/client/:id",controller.UpdateClient)
	r.DELETE("/client/:id",controller.DeleteClient)
	r.Run()
}
