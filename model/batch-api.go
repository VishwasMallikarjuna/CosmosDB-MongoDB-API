package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// EmployeeCreatePayload ...
	BatchCreatePayload struct {
		BATCHID     string `json:"batchId"`
		TENANTID    string `json:"tenantId"`
		RECORDCOUNT int    `json:"recordCount"`
		STATUS      string `json:"status"`
		TOPIC       string `json:"topic"`
	}
)

// ConvertToBSON ....
func (payload BatchCreatePayload) ConvertToBSON() BatchBSON {
	result := BatchBSON{
		ID:          primitive.NewObjectID(),
		BATCHID:     payload.BATCHID,
		TENANTID:    payload.TENANTID,
		RECORDCOUNT: payload.RECORDCOUNT,
		STATUS:      payload.STATUS,
		TOPIC:       payload.TOPIC,
	}
	return result
}

// Validate BatchCreatePayload
func (payload BatchCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.BATCHID,
			validation.Required.Error("batch id is required"),
		),
	)
}
