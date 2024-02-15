package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// EmployeeBSON ...
	EmployeeBSON struct {
		ID         primitive.ObjectID `bson:"_id"`
		EMPID      string             `bson:"empId"`
		Name       string             `bson:"name"`
		DEPARTMENT string             `bson:"department"`
	}

	// EmployeeDetail ...
	EmployeeDetail struct {
		ID         primitive.ObjectID `json:"_id"`
		EMPID      string             `json:"empId"`
		Name       string             `json:"name"`
		DEPARTMENT string             `json:"department"`
	}

	// EmployeeBSON ...
	EmployeeAddressBSON struct {
		ID            primitive.ObjectID `bson:"_id"`
		TEST_SHARD_ID string             `bson:"test_shard_id"`
		EMPID         string             `bson:"empId"`
		ADDRESSID     string             `bson:"addressId"`
		AREA          string             `bson:"area"`
		CITY          string             `bson:"city"`
		STATE         string             `bson:"state"`
		COUNTRY       string             `bson:"country"`
	}

	//Address BSON

	AddressBSON struct {
		AREA    string `bson:"area"`
		CITY    string `bson:"city"`
		STATE   string `bson:"state"`
		COUNTRY string `bson:"country"`
	}
	// EmployeeBSON ...
	EmployeeAddressDetail struct {
		ID            primitive.ObjectID `json:"_id"`
		TEST_SHARD_ID string             `json:"test_shard_id"`
		EMPID         string             `json:"empId"`
		ADDRESSID     string             `json:"addressId"`
		//AddressBSON   AddressBSON        `json:"address"`
		AREA    string `json:"area"`
		CITY    string `json:"city"`
		STATE   string `json:"state"`
		COUNTRY string `json:"country"`
	}

	//Address BSON

	AddressDetail struct {
		AREA    string `json:"area"`
		CITY    string `json:"city"`
		STATE   string `json:"state"`
		COUNTRY string `json:"country"`
	}
)
