package controller

import (
	"net/http"
	"strconv"
	"wms/domain"
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

func (c *InventoryController) UpdateInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var inventory domain.Inventory
		if err := ctx.ShouldBindJSON(&inventory); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		err := c.service.UpdateInventory(ctx, inventory)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Inventory updated successfully"})
	}
}

func (c *InventoryController) ValidateInventory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuID, _ := strconv.Atoi(ctx.Query("sku_id"))
		hubID, _ := strconv.Atoi(ctx.Query("hub_id"))
		quantity, _ := strconv.Atoi(ctx.Query("quantity"))

		if skuID == 0 || hubID == 0 || quantity <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
			return
		}

		isAvailable, err := c.service.ValidateInventory(ctx, skuID, hubID, quantity)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to validate inventory"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"available": isAvailable})
	}
}
