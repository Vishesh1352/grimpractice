package main

import (
	"example.com/grimpractice/controller"
	"example.com/grimpractice/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.DBconnectDatabase()

	//for client table
	r.GET("/client", controller.GetClients)
	r.POST("/client", controller.CreateClient)
	r.GET("/client/:clientID", controller.FindClientById)
	// r.PATCH("/client/:clientID", controller.UpdateClient)
	// r.DELETE("/client/:clientID", controller.DeleteClient)

	//for Employee table
	r.GET("/employee", controller.GetEmployees)
	r.POST("/employee", controller.CreateEmployee)
	r.GET("/employee/:employeeId", controller.FindEmployeeById)
	// r.PATCH("/employee/:employeeId", controller.UpdateEmployee)
	// r.DELETE("/employee/:employeeId", controller.DeleteEmployee)

	//for Work table
	// r.GET("/work", controller.FindEmployeeWorks)
	// r.POST("/work", controller.CreateEmployeeWork)
	// r.GET("/work/employee/:id", controller.FindEmployeeWorkid)
	// r.GET("/work/client/:id", controller.FindClientWorkid)
	// r.PATCH("/work/employee/:id/client/:id/date/:date", controller.UpdateEmployeeWork)
	// r.DELETE("/work/employee/:id", controller.DeleteEmployeeWork)
	// r.DELETE("/work/client/:id", controller.DeleteClientWork)
	// r.DELETE("/work/employee/:id/client/:cid/date/:date", controller.DeleteWork)

	r.POST("/client/:clientId/employee/:employeeId/work", controller.AddWork)
	r.PUT("/client/:clientId/employee/:employeeId/work", controller.UpdateWork)
	//r.DELETE("/client/:clientId/employee/:employeeId/work", controller.DeleteWork)

	r.GET("/bill/client/:clientId", controller.GetBillForAClient)
	r.GET("/bill/employee/:employeeId", controller.GetBillForAnEmployee)

	r.Run()
}
