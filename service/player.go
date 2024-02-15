package service

import (
	"errors"

	"azurepoc/dao"
	"azurepoc/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlayerCreate ...
func PlayerCreate(payload model.PlayerCreatePayload) (model.PlayerBSON, error) {
	var (
		player = payload.ConvertToBSON()
	)

	//Create user
	doc, err := dao.PlayerCreate(player)
	if err != nil {
		err = errors.New("can not create user")
		return doc, err
	}

	return doc, err
}

// PlayerFindByID ...
func PlayerFindByID(playerID primitive.ObjectID) (model.PlayerBSON, error) {
	doc, err := dao.PlayerFindByID(playerID)
	return doc, err
}
