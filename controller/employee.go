package controller

import (
	"fmt"

	"azurepoc/model"
	"azurepoc/service"
	"azurepoc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// EmployeeCreate ...
func EmployeeCreate(c echo.Context) error {
	fmt.Println("EmployeeCreate")
	var (
		payload = c.Get("payload").(model.EmployeeCreatePayload)
	)

	fmt.Println(payload)
	// Process data
	rawData, err := service.EmployeeCreate(payload)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":        rawData.ID,
		"empId":      rawData.EMPID,
		"name":       rawData.Name,
		"department": rawData.DEPARTMENT,
	}, "")
}

// EmployeeCreate ...
func EmployeeAddressCreate(c echo.Context) error {
	fmt.Println("EmployeeAddressCreate ")
	var (
		payload = c.Get("payload").(model.EmployeeAddressCreatePayload)
	)

	fmt.Println(payload)
	// Process data
	rawData, err := service.EmployeeAddressCreate(payload)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	//kafka topic creation

	// Success
	return util.Response200(c, bson.M{
		"_id":           rawData.ID,
		"empId":         rawData.EMPID,
		"addressId":     rawData.ADDRESSID,
		"test_shard_id": rawData.TEST_SHARD_ID,
		"area":          rawData.AREA,
		"city":          rawData.CITY,
		"state":         rawData.STATE,
		"country":       rawData.COUNTRY,
	}, "")
}

// EmployeeFindByID ....
func EmployeeFindByID(c echo.Context) error {
	/*var (
		employer = c.Get("emp").(model.EmployeeBSON)
		empID    = employer.EMPID
	)*/
	empId := c.Param("empId")

	fmt.Println("empId ", empId)
	// Process data
	rawData, err := service.EmployeeFindByID(empId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, rawData, "")
}

// EmployeeFindByID ....
func EmployeeAddressFindByID(c echo.Context) error {
	/*var (
		employer = c.Get("emp").(model.EmployeeBSON)
		empID    = employer.EMPID
	)*/
	addressId := c.QueryParam("addressId")

	fmt.Println("addressId ", addressId)
	// Process data
	rawData, err := service.EmployeeFindByAddressID(addressId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, rawData, "")
}

// update employee by EmpID ....
func EmployeeUpdate(c echo.Context) error {
	fmt.Println("EmployeeUpdate")
	var emp model.EmployeeCreatePayload

	//validate the request body
	if err := c.Bind(&emp); err != nil {
		return err
	}

	fmt.Println("emp", emp)
	// Process data
	rawData, err := service.EmployeeUpdate(emp)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":        rawData.ID,
		"empId":      rawData.EMPID,
		"name":       rawData.Name,
		"department": rawData.DEPARTMENT,
	}, "")
}

// update employee Address by addressId ....
func EmployeeAddressUpdate(c echo.Context) error {
	fmt.Println("EmployeeAddressUpdate")
	var empAddrs model.EmployeeAddressCreatePayload

	var addressId = c.QueryParam("addressId")
	fmt.Println("AddressId to update ", addressId)
	//validate the request body
	if err := c.Bind(&empAddrs); err != nil {
		return err
	}

	fmt.Println("empAddress", empAddrs)
	empAddrs.ADDRESSID = addressId
	// Process data
	rawData, err := service.EmployeeAddressUpdate(empAddrs)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":           rawData.ID,
		"test_shard_id": rawData.TEST_SHARD_ID,
		"empId":         rawData.EMPID,
		"addressId":     addressId,
		"area":          rawData.AREA,
		"city":          rawData.CITY,
		"state":         rawData.STATE,
		"country":       rawData.COUNTRY,
	}, "")
}

// EmployeeDeleteByID ....
func EmployeeDelete(c echo.Context) error {

	empId := c.Param("empId")

	fmt.Println("empId to Delete", empId)
	// Process data
	err := service.EmployeeDelete(empId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	// Success
	return util.Response200(c, bson.M{
		"empId": empId,
	}, "")
}

// EmployeeDeleteByID ....
func EmployeeAddressDelete(c echo.Context) error {

	addressId := c.QueryParam("addressId")

	fmt.Println("Address to Delete", addressId)
	// Process data
	err := service.EmployeeAddressDelete(addressId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	// Success
	return util.Response200(c, bson.M{
		"addressId": addressId,
	}, "")
}
