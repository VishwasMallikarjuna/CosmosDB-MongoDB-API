package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// EmployeeCreatePayload ...
	EmployeeCreatePayload struct {
		EMPID      string `json:"empId"`
		Name       string `json:"name"`
		DEPARTMENT string `json:"department"`
	}
)
type (
	// EmployeeCreatePayload ...
	EmployeeAddressCreatePayload struct {
		//TEST_SHARD_ID string             `json:"test_shard_id"`
		EMPID     string `json:"empId"`
		ADDRESSID string `json:"addressId"`
		AREA      string `json:"area"`
		CITY      string `json:"city"`
		STATE     string `json:"state"`
		COUNTRY   string `json:"country"`
	}
)

// Validate EmployeeCreatePayload
func (payload EmployeeCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		/*validation.Field(
			&payload.EMPID,
			validation.Required.Error("employee id is required"),
		),*/
		validation.Field(
			&payload.Name,
			validation.Required.Error("name is required"),
			validation.Length(3, 30).Error("name is length: 3 -> 30"),
			is.Alpha.Error("name is alpha"),
		),
	)
}

// ConvertToBSON ....
func (payload EmployeeCreatePayload) ConvertToBSON() EmployeeBSON {
	result := EmployeeBSON{
		ID:         primitive.NewObjectID(),
		EMPID:      payload.EMPID,
		Name:       payload.Name,
		DEPARTMENT: payload.DEPARTMENT,
	}
	return result
}

// ConvertToBSON ....
func (payload EmployeeAddressCreatePayload) ConvertToEmployeeAdrBSON() EmployeeAddressBSON {
	result := EmployeeAddressBSON{
		ID:        primitive.NewObjectID(),
		EMPID:     payload.EMPID,
		ADDRESSID: payload.ADDRESSID,
		//TEST_SHARD_ID: payload.TEST_SHARD_ID,
		AREA:    payload.AREA,
		CITY:    payload.CITY,
		STATE:   payload.STATE,
		COUNTRY: payload.COUNTRY,
	}
	return result
}
