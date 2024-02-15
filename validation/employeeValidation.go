package validation

import (
	//"TestGoAPI/dao"

	"azurepoc/model"
	"azurepoc/util"

	"github.com/labstack/echo/v4"
)

// EmployeeCreate ...
func EmployeeCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeCreatePayload
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

// EmployeeCreate ...
func EmployeeAddressCreate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeAddressCreatePayload
		)

		// ValidateStruct
		c.Bind(&payload)
		//err := payload.Validate()

		// Success
		c.Set("payload", payload)
		return next(c)
	}
}

// PlayerValidateID ...
func EmployeeValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			payload model.EmployeeCreatePayload
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

/*func EmployeeCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		var employeeId = c.Param("id")

		param, err := strconv.ParseInt(employeeId, 10, 64)

		if err != nil {
			return nil
		}

		emp, _ := dao.EmployeeFindByID(int32(param))

		// check existed
		if player.ID.IsZero() {
			return util.Response404(c, nil, "Can not find player")
		}

		c.Set("employer", emp)
		return next(c)
	}
}*/
