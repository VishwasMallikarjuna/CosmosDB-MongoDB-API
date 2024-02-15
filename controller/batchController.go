package controller

import (
	"fmt"
	"net/http"

	"azurepoc/config"
	"azurepoc/model"
	"azurepoc/service"
	"azurepoc/util"
	auth "azurepoc/validation"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var envVars = config.GetEnv()

// CreateBatch ...
func CreateBatch(c echo.Context) error {
	fmt.Println("CreateBatch")
	var (
		payload = c.Get("payload").(model.BatchCreatePayload)
	)

	//Add Azure AD authentication
	requestId := c.Response().Header().Get(echo.HeaderXRequestID)
	fmt.Println("RequestId ", requestId)

	//Add authentication and authorization - Azure AD.

	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), payload.TENANTID)
		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		claims.Roles = append(claims.Roles, config.HriIntegrator) //TODO remove after roles added

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail("401", msg))
		}

		if !claims.HasScope(config.HriIntegrator) {
			msg := fmt.Sprintf(config.MsgIntegratorRoleRequired, "create")
			fmt.Println(msg)
			//return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(requestId, msg))
			return nil //TODO - added to skip roles validation
		}

		fmt.Println("envVars.KafkaDisabled ", envVars.KafkaDisabled)
		// if envVars.KafkaDisabled == "false" {
		// 	//write messages to kafka topic
		// }
	}

	fmt.Println(payload)
	// Process data
	rawData, err := service.CreateBatch(payload)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":         rawData.ID,
		"batchId":     rawData.BATCHID,
		"tenantId":    rawData.TENANTID,
		"recordcount": rawData.RECORDCOUNT,
		"status":      rawData.STATUS,
		"topic":       rawData.TOPIC,
	}, "")
}

// BatchFindByBatchID ....
func BatchFindByBatchID(c echo.Context) error {
	batchId := c.QueryParam("batchId")

	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), "")

		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(batchId, msg))
		}

	}
	fmt.Println("batchId ", batchId)
	// Process data
	rawData, err := service.BatchDetailsFindByBatchID(batchId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, rawData, "")
}

// BatchFindByBatchID ....
func BatchFindByTenantID(c echo.Context) error {
	tenantId := c.QueryParam("tenantId")
	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), tenantId)

		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(tenantId, msg))
		}

	}

	fmt.Println("BatchFindByTenantID ", tenantId)
	// Process data
	rawData, err := service.BatchDetailsFindByTenantID(tenantId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	return util.Response200(c, rawData, "")
}

// update batchdetails by batchId ....
func BatchUpdate(c echo.Context) error {
	fmt.Println("BatchUpdate")
	var batch model.BatchCreatePayload

	batchIdToUpdate := c.QueryParam("batchId")
	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), "")

		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(batchIdToUpdate, msg))
		}

	}

	fmt.Println("BatchId to update BatchDetails ", batchIdToUpdate)

	//validate the request body
	if err := c.Bind(&batch); err != nil {
		return err
	}
	// Process data
	rawData, err := service.BatchUpdate(batch)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":         rawData.ID,
		"batchId":     rawData.BATCHID,
		"tenantId":    rawData.TENANTID,
		"recordCount": rawData.RECORDCOUNT,
		"status":      rawData.STATUS,
		"topic":       rawData.TOPIC,
	}, "")
}

// update batchdetails by status ....
func UpdateBatchStatus(c echo.Context) error {
	fmt.Println("UpdateBatchStatus")

	statusToBeUpdated := c.QueryParam("status")
	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), "")

		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail("", msg))
		}

	}

	fmt.Println("statusToBeUpdated ", statusToBeUpdated)
	var batch model.BatchCreatePayload

	//validate the request body
	if err := c.Bind(&batch); err != nil {
		return err
	}
	// Process data
	rawData, err := service.UpdateBatchStatus(batch, statusToBeUpdated)

	fmt.Println(rawData)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{
		"_id":         rawData.ID,
		"batchId":     rawData.BATCHID,
		"tenantId":    rawData.TENANTID,
		"recordCount": rawData.RECORDCOUNT,
		"status":      rawData.STATUS,
		"topic":       rawData.TOPIC,
	}, "")
}

// BatchDeleteBy batchID ....
func BatchDeleteByBatchId(c echo.Context) error {

	batchId := c.QueryParam("batchId")

	if envVars.JWTConfig.AuthDisabled == "false" {
		jwtValidator := auth.NewValidator(envVars.JWTConfig.OidcIssuer, envVars.JWTConfig.JwtAudience)
		//JWT claims validation
		claims, errResp := jwtValidator.GetValidatedClaims("",
			c.Request().Header.Get(echo.HeaderAuthorization), "")

		if errResp != nil {
			return c.JSON(errResp.Code, errResp.Body)
		}

		// validate that the Subject claim (integrator ID) is not missing
		if claims.Subject == "" {
			msg := fmt.Sprintf(config.MsgSubClaimRequiredInJwt)
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(batchId, msg))
		}

		// validate that caller has sufficient permissions
		if !claims.HasScope(auth.HriIntegrator) {
			msg := fmt.Sprintf(auth.MsgIntegratorRoleRequired, "create")
			fmt.Println(msg)
			return c.JSON(http.StatusUnauthorized, model.NewErrorDetail(batchId, msg))
		}

	}
	fmt.Println("batchDetails to Delete", batchId)
	// Process data
	err := service.BatchDeleteByBatchId(batchId)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success
	// Success
	return util.Response200(c, bson.M{
		"batchId": batchId,
	}, "")
}
