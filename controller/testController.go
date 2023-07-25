package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"testing-echo/domain"
	"testing-echo/model"
	"testing-echo/payload"
	"testing-echo/service"
)

type TestController struct {
	testService service.TestService
}

func NewTestController(service service.TestService) *TestController {
	return &TestController{
		testService: service,
	}
}

func Test(c echo.Context) error {

	return c.String(http.StatusOK, "hello")
}

func (controller *TestController) Case1(context echo.Context) error {
	var response payload.GreetingResponse

	response.Greeting = "hello"

	return context.JSON(http.StatusOK, &response)
}

func (controller *TestController) Case2(context echo.Context) error {
	body := payload.GreetingRequest{}

	if err := context.Bind(&body); err != nil {
		return err
	}

	var response payload.GreetingResponse

	response.Greeting = "hello " + body.Name

	return context.JSON(http.StatusOK, &response)
}

func (controller *TestController) Case3(context echo.Context) error {
	numberString := context.Param("number")
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		return context.String(http.StatusInternalServerError, "error")
	}
	result := controller.testService.CalculateFibonacci(number)

	var response payload.CalculateFibonacciResponse

	response.Number = result

	return context.JSON(http.StatusOK, &response)

}

func (controller *TestController) Case4(context echo.Context) error {
	body := payload.PersonRequestBody{}

	if err := context.Bind(&body); err != nil {
		return err
	}

	var person model.Person

	person.FirstName = body.FirstName
	person.LastName = body.LastName
	person.YearOfBirth = body.YearOfBirth

	person, err := controller.testService.CreateGetDeletePersonTestCase(person)
	if err != nil {
		return err
	}

	var personResponse payload.PersonResponseBody

	personResponse.Id = person.Id
	personResponse.FirstName = person.FirstName
	personResponse.LastName = person.LastName
	personResponse.YearOfBirth = person.YearOfBirth

	return context.JSON(http.StatusOK, &personResponse)
}

func (controller *TestController) Case5(context echo.Context) error {
	body := payload.PersonRequestBody{}

	if err := context.Bind(&body); err != nil {
		return err
	}

	var person domain.Person

	person.FirstName = body.FirstName
	person.LastName = body.LastName
	person.YearOfBirth = body.YearOfBirth

	person, err := controller.testService.CreateGetDeletePersonORMTestCase(person)
	if err != nil {
		return err
	}

	var personResponse payload.PersonResponseBody

	personResponse.Id = person.Id
	personResponse.FirstName = person.FirstName
	personResponse.LastName = person.LastName
	personResponse.YearOfBirth = person.YearOfBirth

	return context.JSON(http.StatusOK, &personResponse)
}
