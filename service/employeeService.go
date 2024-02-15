package service

import (
	"errors"
	"fmt"

	"azurepoc/dao"
	"azurepoc/model"
)

const (
	test_shard_id = "SH123"
)

// EmployeeCreate ...
func EmployeeCreate(payload model.EmployeeCreatePayload) (model.EmployeeBSON, error) {
	var (
		employee = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.EmployeeCreate(employee)
	if err != nil {
		err = errors.New("failed to create Employee details")
		return doc, err
	}

	return doc, err
}

// EmployeeCreate ...
func EmployeeAddressCreate(payload model.EmployeeAddressCreatePayload) (model.EmployeeAddressBSON, error) {
	var (
		employee = payload.ConvertToEmployeeAdrBSON()
	)

	employee.TEST_SHARD_ID = test_shard_id
	//Create user
	doc, err := dao.EmployeeAddressCreate(employee)
	if err != nil {
		err = errors.New("failed to create Employee details")
		return doc, err
	}

	return doc, err
}

// EmployeeFindByID ...
func EmployeeFindByID(empId string) (model.EmployeeBSON, error) {
	doc, err := dao.EmployeeFindByID(empId)
	return doc, err
}

// EmployeeFindByID ...
func EmployeeFindByAddressID(addressId string) (model.EmployeeAddressBSON, error) {
	fmt.Println("EmployeeFindByAddressID ", addressId)
	doc, err := dao.EmployeeFindByAddressID(addressId)
	return doc, err
}

//Update Employee

func EmployeeUpdate(payload model.EmployeeCreatePayload) (model.EmployeeBSON, error) {
	//func EmployeeUpdate(payload model.EmployeeCreatePayload) (model.EmployeeBSON, error) {
	var (
		employee = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.EmployeeUpdate(employee)
	if err != nil {
		err = errors.New("failed to create Employee details")
		return doc, err
	}

	return doc, err
}

//Update Employee

func EmployeeAddressUpdate(payload model.EmployeeAddressCreatePayload) (model.EmployeeAddressBSON, error) {
	//func EmployeeUpdate(payload model.EmployeeCreatePayload) (model.EmployeeBSON, error) {
	var (
		employeeAddrs = payload.ConvertToEmployeeAdrBSON()
	)

	//Create user
	doc, err := dao.EmployeeAddressUpdate(employeeAddrs)
	if err != nil {
		err = errors.New("failed to create Employee details")
		return doc, err
	}

	return doc, err
}
func EmployeeDelete(employeeId string) error {
	err := dao.EmployeeDelete(employeeId)
	return err
}

func EmployeeAddressDelete(addressId string) error {
	err := dao.EmployeeAddressDelete(addressId)
	return err
}
