package controller

import (
	"net/http"
	"strconv"
	"wms/logger"
	"wms/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type InventoryController struct {
	service service.InventoryService
}

func NewInventoryController(s service.InventoryService) *InventoryController {
	return &InventoryController{
		service: s,
	}
}

// GetInventoryDetails retrieves inventory based on hub_id and sku_id
func (c *InventoryController) GetInventoryDetails() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubID, err := strconv.Atoi(ctx.Param("hub_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
			return
		}

		skuID, err := strconv.Atoi(ctx.Param("sku_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid SKU ID"})
			return
		}

		inventory, err := c.service.FetchInventory(ctx, hubID, skuID)
		if err != nil {
			logger.Error("Failed to fetch inventory", logrus.Fields{"hub_id": hubID, "sku_id": skuID, "error": err.Error()})

			if err.Error() == "record not found" {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Inventory not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving inventory"})
			}
			return
		}

		ctx.JSON(http.StatusOK, inventory)
	}
}
