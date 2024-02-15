package validation

import (
	//"TestGoAPI/dao"

	"azurepoc/model"
	"azurepoc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeCreate ...
func CreateBatch(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.BatchCreatePayload
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
