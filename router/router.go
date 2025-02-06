package router

import (
	"context"
	"wms/controllers"

	"github.com/omniful/go_commons/http"
)

func InternalRoutes(ctx context.Context, server *http.Server) (err error) {
	router := server.Engine // Using Gin Engine from server

	// API Versioning (optional)
	api := router.Group("/api/v1")
	{
		// Warehouse Hub Routes
		api.GET("/hubs/:id", controllers.GetHub) // View a single hub
		api.POST("/hubs", controllers.CreateHub) // Create a new hub

		// SKU Routes
		api.GET("/skus", controllers.GetSKUs) // Get multiple SKUs
		api.POST("/skus", controllers.CreateSKU)

		// Inventory Routes
		api.GET("/inventory", controllers.GetInventory)  // View inventory
		api.PUT("/inventory", controllers.EditInventory) // Edit inventory
	}
	return nil
}
