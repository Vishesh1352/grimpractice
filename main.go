package main

import (
	"github.com/gin-gonic/gin"

	"example.com/grimpractice/controller"
	"example.com/grimpractice/models"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()

	//for client table
	r.GET("/client", controller.GetClients)
	r.POST("/client", controller.CreateClient)
	r.GET("/client/:clientID", controller.FindClientById)
	r.PATCH("/client/:clientID", controller.UpdateClient)
	r.DELETE("/client/:clientID", controller.DeleteClient)

	//for Employee table
	r.GET("/employee", controller.GetEmployees)
	r.POST("/employee", controller.CreateEmployee)
	r.GET("/employee/:employeeId", controller.FindEmployeeById)
	r.PATCH("/employee/:employeeId", controller.UpdateEmployee)
	r.DELETE("employee/:employeeId", controller.DeleteEmployee)

	//for Work
	r.POST("/client/:clientId/employee/:employeeId/work", controller.AddWork)
	r.PUT("/client/:clientId/employee/:employeeId/work", controller.UpdateWork)

	//for Billing
	r.GET("/bill/client/:clientId", controller.GetBillForAClient)
	r.GET("/bill/employee/:employeeId", controller.GetBillForAnEmployee)

	r.Run()
}
