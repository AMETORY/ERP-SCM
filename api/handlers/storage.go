package handlers

import (
	"sample-scm-backend/objects"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/distribution"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	ctx             *context.ERPContext
	distributionSrv *distribution.DistributionService
}

func NewStorageHandler(ctx *context.ERPContext) *StorageHandler {
	distributionSrv, ok := ctx.DistributionService.(*distribution.DistributionService)
	if !ok {
		panic("distribution service is not found")
	}
	return &StorageHandler{ctx: ctx, distributionSrv: distributionSrv}
}

func (s *StorageHandler) CreateWarehouseHandler(c *gin.Context) {
	var input objects.WarehouseRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	warehouse := models.WarehouseModel{
		Name:    input.Name,
		Address: input.Address,
	}
	err = s.distributionSrv.StorageService.CreateWarehouse(&warehouse, input.Lat, input.Lng)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Warehouse created successfully", "data": warehouse})
}

func (s *StorageHandler) UpdateWarehouseHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.WarehouseModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = s.distributionSrv.StorageService.UpdateWarehouse(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Warehouse updated successfully"})
}

func (s *StorageHandler) DeleteWarehouseHandler(c *gin.Context) {
	id := c.Param("id")

	err := s.distributionSrv.StorageService.DeleteWarehouse(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Warehouse deleted successfully"})
}

func (s *StorageHandler) CreateWarehouseLocationHandler(c *gin.Context) {
	var input objects.LocationPointRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	location := models.LocationPointModel{
		Name:      input.Name,
		Address:   input.Address,
		Latitude:  input.Lat,
		Longitude: input.Lng,
		Type:      input.Type,
	}

	var warehouse *models.WarehouseModel
	if input.Type == "WAREHOUSE" {
		warehouse = &models.WarehouseModel{
			Name:    input.Name,
			Address: input.Address,
		}
	}
	err = s.distributionSrv.StorageService.CreateLocation(&location, warehouse)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	location.Warehouse = warehouse
	c.JSON(201, gin.H{"message": "WarehouseLocation created successfully", "data": location})
}

func (s *StorageHandler) UpdateWarehouseLocationHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.LocationPointModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = s.distributionSrv.StorageService.GetWarehouseLocationByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = s.distributionSrv.StorageService.UpdateWarehouseLocation(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "WarehouseLocation updated successfully"})
}

func (s *StorageHandler) DeleteWarehouseLocationHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := s.distributionSrv.StorageService.GetWarehouseLocationByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = s.distributionSrv.StorageService.DeleteWarehouseLocation(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "WarehouseLocation deleted successfully"})
}

func (s *StorageHandler) GetWarehouseLocationByIDHandler(c *gin.Context) {
	id := c.Param("id")
	data, err := s.distributionSrv.StorageService.GetWarehouseLocationByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "WarehouseLocation retrieved successfully", "data": data})
}

func (s *StorageHandler) GetWarehouseLocationsHandler(c *gin.Context) {
	search, _ := c.GetQuery("search")
	data, err := s.distributionSrv.StorageService.GetWarehouseLocations(*c.Request, search)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "WarehouseLocations retrieved successfully", "data": data})
}
