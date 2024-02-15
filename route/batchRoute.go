package route

import (
	"azurepoc/controller"
	"azurepoc/validation"

	"github.com/labstack/echo/v4"
)

func batch(e *echo.Echo) {
	batch := e.Group("/batch")
	batch.POST("", controller.CreateBatch, validation.CreateBatch)
	batch.GET("/tenant", controller.BatchFindByTenantID)
	batch.GET("", controller.BatchFindByBatchID)
	batch.PUT("", controller.BatchUpdate)
	batch.PUT("/update", controller.UpdateBatchStatus)
	batch.DELETE("", controller.BatchDeleteByBatchId)
}
