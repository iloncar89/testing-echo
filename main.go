package main

import (
	"fmt"
	"testing-echo/config"
	"testing-echo/controller"
	"testing-echo/dao"
	"testing-echo/router"
	"testing-echo/service"
)

func main() {
	// Database ORM
	dbOrm := config.DatabaseConnectionOrm()

	// Database Plain
	db := config.DatabaseConnection()

	// DAO
	testDao := dao.NewPersonDaoImpl(db, dbOrm)

	// Service
	testService := service.NewTestServiceImpl(testDao)

	// Controller
	testController := controller.NewTestController(testService)

	// Echo instance
	routes := router.NewRouter(testController)

	routes.Logger.Fatal(routes.Start(":8083"))

	fmt.Println("Server running on port 8083")
}
