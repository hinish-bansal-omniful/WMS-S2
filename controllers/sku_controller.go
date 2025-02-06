package controller

import (
	"net/http"
	"strconv"
	"wms/domain"
	"wms/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type SkuController struct {
	service service.SKUService
}

func NewSkuController(s service.SKUService) *SkuController {
	return &SkuController{service: s}
}

func (c *SkuController) GetSKU() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		skuID := ctx.Param("sku_id")

		sku, err := c.service.GetSKUByID(ctx, skuID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"sku_id": skuID,
				"error":  err.Error(),
			}).Error("Failed to fetch SKU")

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SKU"})
			return
		}

		if sku == nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "SKU not found"})
			return
		}

		ctx.JSON(http.StatusOK, sku)
	}
}

func (c *SkuController) GetSkuBySellerID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellerID, err := strconv.Atoi(ctx.Param("seller_id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Seller ID"})
			return
		}

		skus, err := c.service.FetchSkuBySellerID(ctx, sellerID)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"seller_id": sellerID,
				"error":     err.Error(),
			}).Error("Failed to fetch SKUs for seller")

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SKUs"})
			return
		}

		if len(skus) == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No SKUs found for the seller"})
			return
		}

		ctx.JSON(http.StatusOK, skus)
	}
}

func (c *SkuController) CreateSKU() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sku domain.SKU
		if err := ctx.ShouldBindJSON(&sku); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		createdSKU, err := c.service.CreateSKU(ctx, &sku)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("Failed to create SKU")

			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create SKU"})
			return
		}

		ctx.JSON(http.StatusCreated, createdSKU)
	}
}
