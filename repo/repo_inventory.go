package repo

import (
	"context"
	"errors"
	"wms/domain"
	"wms/logger"

	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	GetInventoryByHubAndSKU(ctx context.Context, hubID int, skuID int) (domain.Inventory, error)
}

type inventoryRepository struct {
	db *postgres.DbCluster
}

func NewInventoryRepository(db *postgres.DbCluster) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetInventoryByHubAndSKU(ctx context.Context, hubID int, skuID int) (domain.Inventory, error) {
	var inventory domain.Inventory

	result := r.db.GetMasterDB(ctx).Where("hub_id = ? AND sku_id = ?", hubID, skuID).First(&inventory)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Inventory{}, errors.New("record not found")
		}
		logger.Error("Error retrieving inventory from database", logrus.Fields{"hub_id": hubID, "sku_id": skuID, "error": result.Error.Error()})
		return domain.Inventory{}, result.Error
	}

	return inventory, nil
}
