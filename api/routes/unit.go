package routes

import (
	"sample-scm-backend/api/handlers"
	"sample-scm-backend/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupUnitRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	unitHandler := handlers.NewUnitHandler(erpContext)
	unitGroup := r.Group("/unit")
	unitGroup.Use(middlewares.AuthMiddleware(erpContext, false))
	{
		unitGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:unit:read"}), unitHandler.ListUnitsHandler)
		unitGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:unit:read"}), unitHandler.GetUnitHandler)
		unitGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:unit:create"}), unitHandler.CreateUnitHandler)
		unitGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:unit:update"}), unitHandler.UpdateUnitHandler)
		unitGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:unit:delete"}), unitHandler.DeleteUnitHandler)
	}
}
