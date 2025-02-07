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
	UpdateInventory(ctx context.Context, inventory domain.Inventory) error
	ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error)
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

func (r *inventoryRepository) UpdateInventory(ctx context.Context, inventory domain.Inventory) error {
	result := r.db.GetMasterDB(ctx).Model(&inventory).
		Where(" hub_id = ? AND sku_id = ?", inventory.HubID, inventory.SKUID).
		Updates(inventory)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no inventory record found to update")
	}

	return nil
}
func (r *inventoryRepository) ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error) {
	var totalQuantity int
	result := r.db.GetMasterDB(ctx).Table("inventories").
		Where("sku_id = ? AND hub_id = ?", skuID, hubID).
		Select("SUM(quantity)").Row().Scan(&totalQuantity)

	if result != nil {
		return false, errors.New("error fetching inventory data")
	}

	return totalQuantity >= quantity, nil
}
