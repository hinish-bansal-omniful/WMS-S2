package router

import (
	"context"
	controller "wms/controllers"

	postgres "wms/db"
	"wms/repo"
	"wms/service"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/http"
)

func InternalRoutes(ctx context.Context, s *http.Server) (err error) {
	rtr := s.Engine.Group("/api/v1")

	// todo use go wire if needed
	newRepository := repo.NewRepository(postgres.GetCluster().DbCluster)
	newService := service.NewService(newRepository)
	HubController := controller.NewHubController(newService)

	skuRepo := repo.NewSKURepository(postgres.GetCluster().DbCluster)
	skuService := service.NewSKUService(skuRepo)
	skuController := controller.NewSkuController(skuService)

	inventoryRepo := repo.NewInventoryRepository(postgres.GetCluster().DbCluster)
	inventoryService := service.NewInventoryService(inventoryRepo)
	inventoryController := controller.NewInventoryController(inventoryService)

	// make apis for it
	rtr.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "mst"})
	})

	rtr.GET("/get_hubs", HubController.GetHubs())
	rtr.GET("/get_hub/:id", HubController.GetHubByID())
	rtr.POST("/hubs", HubController.CreateHub())

	rtr.GET("/sku/:sku_id", skuController.GetSKU())
	rtr.GET("/sku/seller/:seller_id", skuController.GetSkuBySellerID())
	rtr.POST("/sku", skuController.CreateSKU())

	rtr.GET("/inventory/:hub_id/:sku_id", inventoryController.GetInventoryDetails())
	rtr.PUT("/inventory/", inventoryController.UpdateInventory())
	rtr.GET("/inventory/validate", inventoryController.ValidateInventory())

	return
}
