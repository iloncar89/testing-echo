package router

import (
	"testing-echo/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(testController *controller.TestController) *echo.Echo {
	router := echo.New()

	baseRouter := router.Group("/api/v1")
	testRouter := baseRouter.Group("/test")
	testRouter.GET("/case1", testController.Case1)
	testRouter.POST("/case2", testController.Case2)
	testRouter.GET("/case3/:number", testController.Case3)
	testRouter.POST("/case4", testController.Case4)
	testRouter.POST("/case5", testController.Case5)

	return router
}
