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
		//for client table
	r.GET("/client",controller.FindClients)
	r.POST("/client",controller.CreateClients)
	r.GET("/client/:id",controller.FindClientsid)
	r.PATCH("/client/:id",controller.UpdateClient)
	r.DELETE("/client/:id",controller.DeleteClient)

		//for Employee table
	r.GET("/Employee",controller.FindEmployees)
	r.POST("/Employee",controller.CreateEmployees)
	r.GET("/Employee/:id",controller.FindEmployeesid)
	r.PATCH("/Employee/:id",controller.UpdateEmployee)
	r.DELETE("/Employee/:id",controller.DeleteEmployee)

	//for EmployeeWork table
	r.GET("/Worktable",controller.FindEmployeeWorks)
	r.POST("/Worktable",controller.CreateEmployeeWork)
	r.GET("/Worktable/employee/:id",controller.FindEmployeeWorkid)
	r.GET("/Worktable/client/:id",controller.FindClientWorkid)
	r.PATCH("/Worktable/employee/:id/client/:id/date/:date",controller.UpdateEmployeeWork)
	r.DELETE("/Worktable/employee/:id",controller.DeleteEmployeeWork)
	r.DELETE("/Worktable/client/:id",controller.DeleteClientWork)
	r.DELETE("/Worktable/employee/:id/client/:cid/date/:date",controller.DeleteWork)
	r.Run()
}
