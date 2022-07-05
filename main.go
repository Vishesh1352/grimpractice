package main

import (
	"example.com/grimpractice/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.connecToDatabase()

	//for client table
	r.GET("/client", controller.GetClients)
	r.POST("/client", controller.CreateClient)
	r.GET("/client/:clientID", controller.FindClientById)

	//for Employee table
	r.GET("/employee", controller.GetEmployees)
	r.POST("/employee", controller.CreateEmployee)
	r.GET("/employee/:employeeId", controller.FindEmployeeById)

	//for Work
	r.POST("/client/:clientId/employee/:employeeId/work", controller.AddWork)
	r.PUT("/client/:clientId/employee/:employeeId/work", controller.UpdateWork)

	//for Billing
	r.GET("/bill/client/:clientId", controller.GetBillForAClient)
	r.GET("/bill/employee/:employeeId", controller.GetBillForAnEmployee)

	r.Run()
}
