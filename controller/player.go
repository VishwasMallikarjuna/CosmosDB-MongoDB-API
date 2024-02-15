package controller

import (
	"azurepoc/model"
	"azurepoc/service"
	"azurepoc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// PlayerCreate ...
func PlayerCreate(c echo.Context) error {
	var (
		payload = c.Get("payload").(model.PlayerCreatePayload)
	)

	// Process data
	rawData, err := service.PlayerCreate(payload)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")
}

// PlayerFindByID ....
func PlayerFindByID(c echo.Context) error {
	var (
		player   = c.Get("player").(model.PlayerBSON)
		playerID = player.ID
	)

	// Process data
	rawData, err := service.PlayerFindByID(playerID)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, rawData, "")
}
