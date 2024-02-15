package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// BatchBSON ...
	BatchBSON struct {
		ID          primitive.ObjectID `bson:"_id"`
		BATCHID     string             `bson:"batchId"`
		TENANTID    string             `bson:"tenantId"`
		RECORDCOUNT int                `bson:"recordCount"`
		STATUS      string             `bson:"status"`
		TOPIC       string             `bson:"topic"`
	}

	// BatchDetail ...
	BatchDetail struct {
		ID          primitive.ObjectID `json:"_id"`
		BATCHID     string             `json:"batchId"`
		TENANTID    string             `json:"tenantId"`
		RECORDCOUNT int                `json:"recordCount"`
		STATUS      string             `json:"status"`
		TOPIC       string             `json:"topic"`
	}
)
