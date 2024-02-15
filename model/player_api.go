package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// PlayerCreatePayload ...
	PlayerCreatePayload struct {
		Name string `json:"name"`
	}
)

// Validate PlayerCreatePayload
func (payload PlayerCreatePayload) Validate() error {
	return validation.ValidateStruct(&payload,
		validation.Field(
			&payload.Name,
			validation.Required.Error("name is required"),
			validation.Length(3, 30).Error("name is length: 3 -> 30"),
			is.Alpha.Error("name is alpha"),
		),
	)
}

// ConvertToBSON ....
func (payload PlayerCreatePayload) ConvertToBSON() PlayerBSON {
	result := PlayerBSON{
		ID:        primitive.NewObjectID(),
		Name:      payload.Name,
		CreatedAt: time.Now(),
	}
	return result
}
