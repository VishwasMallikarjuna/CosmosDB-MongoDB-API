package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	// PlayerBSON ...
	PlayerBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		Name      string             `bson:"name"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	// PlayerDetail ...
	PlayerDetail struct {
		ID        primitive.ObjectID `json:"_id"`
		Name      string             `json:"name"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}
)
