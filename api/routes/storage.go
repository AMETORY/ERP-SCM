package routes

import (
	"sample-scm-backend/api/handlers"
	"sample-scm-backend/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupStorageRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewStorageHandler(erpContext)
	locationPointGroup := r.Group("/storage")
	locationPointGroup.Use(middlewares.AuthMiddleware(erpContext, false))
	{
		locationPointGroup.POST("/warehouse/create", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:create-warehouse"}), handler.CreateWarehouseHandler)
		locationPointGroup.PUT("/warehouse/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:delete-warehouse"}), handler.UpdateWarehouseHandler)
		locationPointGroup.GET("/warehouse/list", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:list-warehouse"}), handler.GetWarehousesHandler)
		locationPointGroup.DELETE("/warehouse/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:delete-warehouse"}), handler.DeleteWarehouseHandler)
		locationPointGroup.POST("/location/create", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:create-create"}), handler.CreateWarehouseLocationHandler)
		locationPointGroup.PUT("/location/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:update-location"}), handler.UpdateWarehouseLocationHandler)
		locationPointGroup.DELETE("/location/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:delete-location"}), handler.DeleteWarehouseLocationHandler)
		locationPointGroup.GET("/location/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:get-location-detail"}), handler.GetWarehouseLocationByIDHandler)
		locationPointGroup.GET("location/list", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:storage:get-locations"}), handler.GetWarehouseLocationsHandler)
	}
}
