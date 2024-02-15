package route

import (
	"azurepoc/controller"
	"azurepoc/validation"

	"github.com/labstack/echo/v4"
)

func employee(e *echo.Echo) {
	emp := e.Group("/employee")
	emp.POST("", controller.EmployeeCreate, validation.EmployeeCreate)
	emp.GET("/:empId", controller.EmployeeFindByID)
	emp.PUT("/:empId", controller.EmployeeUpdate)
	emp.DELETE("/:empId", controller.EmployeeDelete)
	//EmployeeAddress
	emp.POST("/address", controller.EmployeeAddressCreate, validation.EmployeeAddressCreate)
	emp.GET("/address", controller.EmployeeAddressFindByID)
	emp.PUT("/address", controller.EmployeeAddressUpdate)
	emp.DELETE("/address", controller.EmployeeAddressDelete)
}
