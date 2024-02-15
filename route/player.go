package route

import (
	"azurepoc/controller"
	"azurepoc/validation"

	"github.com/labstack/echo/v4"
)

func player(e *echo.Echo) {
	players := e.Group("/players")
	players.POST("", controller.PlayerCreate, validation.PlayerCreate)
	players.GET("/:id", controller.PlayerFindByID, validation.PlayerValidateID, validation.PlayerCheckExistedByID)
}
