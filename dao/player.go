package dao

import (
	"context"

	"azurepoc/model"
	"azurepoc/module/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlayerCreate ...
func PlayerCreate(doc model.PlayerBSON) (model.PlayerBSON, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
	)

	// Insert one
	_, err := userCol.InsertOne(ctx, doc)
	return doc, err
}

// PlayerFindByID ...
func PlayerFindByID(id primitive.ObjectID) (model.PlayerBSON, error) {
	var (
		userCol = database.PlayerCol()
		ctx     = context.Background()
		result  model.PlayerBSON
		filter  = bson.M{"_id": id}
	)

	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}
