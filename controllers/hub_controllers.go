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

type HubController struct {
	service service.HubService
}

func NewHubController(s service.HubService) *HubController {
	return &HubController{
		service: s,
	}
}

func (c *HubController) GetHubs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubs, err := c.service.FetchHubs(ctx)
		if err != nil {
			logger.Error("Failed to fetch hubs", logrus.Fields{"error": err.Error()})
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch hubs",
			})
			return
		}
		ctx.JSON(http.StatusOK, hubs)
	}
}

func (c *HubController) GetHubByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hubID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub ID"})
			return
		}

		hub, err := c.service.FetchHubByID(ctx, hubID)
		if err != nil {
			logger.Error("Hub not found", logrus.Fields{"hub_id": hubID, "error": err.Error()})

			// Differentiate between "not found" and "server errors"
			if err.Error() == "record not found" {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Hub not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving hub"})
			}
			return
		}

		ctx.JSON(http.StatusOK, hub)
	}
}

// **New: Create a Hub API**
func (c *HubController) CreateHub() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hub domain.Hub

		if err := ctx.ShouldBindJSON(&hub); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		createdHub, err := c.service.CreateHub(ctx, hub)
		if err != nil {
			logger.Error("Failed to create hub", logrus.Fields{"error": err.Error()})
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create hub"})
			return
		}

		ctx.JSON(http.StatusCreated, createdHub)
	}
}
