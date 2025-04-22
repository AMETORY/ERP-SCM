package routes

import (
	"sample-scm-backend/api/handlers"
	"sample-scm-backend/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupLogisticRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewLogisticHandler(erpContext)
	logisticGroup := r.Group("/logistic")
	logisticGroup.Use(middlewares.AuthMiddleware(erpContext, false))
	{
		logisticGroup.POST("/create-distribution-event", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:create-distribution-event"}), handler.CreateDistributionEventHandler)
		logisticGroup.GET("/list-distribution-event", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:list-distribution-event"}), handler.ListDistributionEventsHandler)
		logisticGroup.GET("/distribution-event/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:read-distribution-event"}), handler.ReadDistributionEventHandler)
		logisticGroup.GET("/distribution-event/:id/report", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:read-distribution-event"}), handler.GetDistributionEventReportHandler)
		logisticGroup.DELETE("/distribution-event/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:delete-distribution-event"}), handler.DeleteDistributionEventHandler)
		logisticGroup.POST("/create-shipment", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:create-shipment"}), handler.CreateShipmentHandler)
		logisticGroup.GET("/shipment/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:read-shipment"}), handler.ReadShipmentHandler)
		logisticGroup.PUT("/shipment/:id/update-status", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:update-shipment"}), handler.UpdateStatusShipmentHandler)
		logisticGroup.PUT("/shipment/:id/add-item", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:update-shipment"}), handler.AddItemShipmentHandler)
		logisticGroup.DELETE("/shipment/:id/delete-item/:itemId", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:update-shipment"}), handler.DeleteItemShipmentHandler)
		logisticGroup.DELETE("/delete-shipment/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:delete-shipment"}), handler.DeleteShipmentHandler)
		logisticGroup.PUT("/ready-to-ship/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:ready-to-ship"}), handler.ReadyToShipHandler)
		logisticGroup.PUT("/process-shipment/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:process-shipment"}), handler.ProcessShipmentHandler)
		logisticGroup.POST("/create-shipment-leg", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:create-shipment-leg"}), handler.CreateShipmentLegHandler)
		logisticGroup.PUT("/start-shipment-leg/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:start-shipment-leg"}), handler.StartShipmentLegHandler)
		logisticGroup.PUT("/arrived-shipment-leg/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:arrived-shipment-leg"}), handler.ArrivedShipmentLegHandler)
		logisticGroup.PUT("/add-tracking-event/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:add-tracking-event"}), handler.AddTrackingEventHandler)
		logisticGroup.GET("/generate-shipment-report/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:generate-shipment-report"}), handler.GenerateShipmentReportHandler)
		logisticGroup.GET("/generate-distributor-event-report/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:generate-distributor-event-report"}), handler.GenerateDistributorEventReportHandler)
		logisticGroup.POST("/report-lost-damage/:id/:legId", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:report-lost-damage"}), handler.ReportLostDamageHandler)
	}
}
