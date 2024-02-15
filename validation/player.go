package validation

import (
	"fmt"

	"azurepoc/dao"
	"azurepoc/model"
	"azurepoc/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlayerCreate ...
func PlayerCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.PlayerCreatePayload
		)

		// ValidateStruct
		c.Bind(&payload)
		err := payload.Validate()

		//if err
		if err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// Success
		c.Set("payload", payload)
		return next(c)
	}
}

func validationObjectID(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}
	return objectID, err
}

// PlayerValidateID ...
func PlayerValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			id            = c.Param("id")
			playerID, err = validationObjectID(id)
		)

		// if err
		if err != nil {
			return util.Response400(c, nil, "ID not valid")
		}

		c.Set("playerID", playerID)
		return next(c)
	}
}

func PlayerCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			playerID = c.Get("playerID").(primitive.ObjectID)
		)

		player, _ := dao.PlayerFindByID(playerID)

		// check existed
		if player.ID.IsZero() {
			return util.Response404(c, nil, "Can not find player")
		}

		c.Set("player", player)
		return next(c)
	}
}
