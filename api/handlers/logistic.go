package handlers

import (
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/distribution"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type LogisticHandler struct {
	ctx             *context.ERPContext
	distributionSrv *distribution.DistributionService
}

func NewLogisticHandler(ctx *context.ERPContext) *LogisticHandler {
	distributionSrv, ok := ctx.DistributionService.(*distribution.DistributionService)
	if !ok {
		panic("distribution service is not found")
	}

	return &LogisticHandler{ctx: ctx, distributionSrv: distributionSrv}
}

func (h *LogisticHandler) CreateDistributionEventHandler(c *gin.Context) {
	// Implement logic for creating a distribution event
	var input models.DistributionEventModel
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.distributionSrv.LogisticService.CreateDistributionEvent(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Distribution event created successfully"})
}

func (h *LogisticHandler) CreateShipmentHandler(c *gin.Context) {
	var input models.ShipmentModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.distributionSrv.LogisticService.CreateShipment(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment created successfully"})
}

func (h *LogisticHandler) ReadyToShipHandler(c *gin.Context) {
	var input struct {
		Notes *string   `json:"notes"`
		Date  time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := h.distributionSrv.LogisticService.ReadyToShip(id, input.Date, input.Notes)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment marked as ready to ship"})
}

func (h *LogisticHandler) ProcessShipmentHandler(c *gin.Context) {
	var input struct {
		Notes string    `json:"notes"`
		Date  time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := h.distributionSrv.LogisticService.ProcessShipment(id, input.Date, input.Notes)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment processed successfully"})
}

func (h *LogisticHandler) CreateShipmentLegHandler(c *gin.Context) {
	var input models.ShipmentLegModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.distributionSrv.LogisticService.CreateShipmentLeg(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment leg created successfully"})
}

func (h *LogisticHandler) StartShipmentLegHandler(c *gin.Context) {
	var input struct {
		Notes string    `json:"notes"`
		Date  time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := h.distributionSrv.LogisticService.StartShipmentLegDelivery(id, input.Date, input.Notes)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment leg started successfully"})
}

func (h *LogisticHandler) ArrivedShipmentLegHandler(c *gin.Context) {
	var input struct {
		Notes string    `json:"notes"`
		Date  time.Time `json:"date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := h.distributionSrv.LogisticService.ArrivedShipmentLegDelivery(id, input.Date, input.Notes)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Shipment leg marked as arrived"})
}

func (h *LogisticHandler) AddTrackingEventHandler(c *gin.Context) {
	var input struct {
		Status    string  `json:"status"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Notes     string  `json:"notes"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	err := h.distributionSrv.LogisticService.AddTrackingEvent(id, input.Status, input.Latitude, input.Longitude, input.Notes)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Tracking event added successfully"})
}

func (h *LogisticHandler) GenerateShipmentReportHandler(c *gin.Context) {
	id := c.Param("id")
	report, err := h.distributionSrv.LogisticService.GenerateShipmentReport(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": report, "message": "Shipment report generated successfully"})
}

func (h *LogisticHandler) GenerateDistributorEventReportHandler(c *gin.Context) {
	id := c.Param("id")

	var reportID *string
	if c.Query("report_id") != "" {
		repID := c.Query("report_id")
		reportID = &repID
	}
	report, err := h.distributionSrv.LogisticService.GenerateDistributionEventReport(id, reportID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": report, "message": "Distributor event report generated successfully"})
}

func (h *LogisticHandler) ReportLostDamageHandler(c *gin.Context) {
	var input models.IncidentEventModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	legId := c.Param("legId")

	err := h.distributionSrv.LogisticService.ReportLostOrDamage(id, legId, input.OccurredAt, &input, models.MovementType(input.EventType), input.WasteWarehouseID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Lost or damaged goods reported successfully"})
}
