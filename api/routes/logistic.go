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
		logisticGroup.POST("/create-shipment", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:create-shipment"}), handler.CreateShipmentHandler)
		logisticGroup.PUT("/ready-to-ship/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:ready-to-ship"}), handler.ReadyToShipHandler)
		logisticGroup.PUT("/process-shipment/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:process-shipment"}), handler.ProcessShipmentHandler)
		logisticGroup.POST("/create-shipment-leg/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:create-shipment-leg"}), handler.CreateShipmentLegHandler)
		logisticGroup.PUT("/start-shipment-leg/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:start-shipment-leg"}), handler.StartShipmentLegHandler)
		logisticGroup.PUT("/arrived-shipment-leg/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:arrived-shipment-leg"}), handler.ArrivedShipmentLegHandler)
		logisticGroup.POST("/add-tracking-event/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:add-tracking-event"}), handler.AddTrackingEventHandler)
		logisticGroup.GET("/generate-shipment-report/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:generate-shipment-report"}), handler.GenerateShipmentReportHandler)
		logisticGroup.GET("/generate-distributor-event-report/:id", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:generate-distributor-event-report"}), handler.GenerateDistributorEventReportHandler)
		logisticGroup.POST("/report-lost-damage/:id/:legId", middlewares.RbacUserMiddleware(erpContext, []string{"distribution:logistic:report-lost-damage"}), handler.ReportLostDamageHandler)
	}
}
