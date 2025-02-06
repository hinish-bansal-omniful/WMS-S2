package service

import (
	"context"
	"errors"
	"wms/domain"
	"wms/logger"
	"wms/repo"

	"github.com/sirupsen/logrus"
)

type InventoryService interface {
	FetchInventory(ctx context.Context, hubID int, skuID int) (domain.Inventory, error)
}

type InventoryServiceImpl struct {
	repo repo.InventoryRepository
}

func NewInventoryService(r repo.InventoryRepository) InventoryService {
	return &InventoryServiceImpl{repo: r}
}

func (s *InventoryServiceImpl) FetchInventory(ctx context.Context, hubID int, skuID int) (domain.Inventory, error) {
	inventory, err := s.repo.GetInventoryByHubAndSKU(ctx, hubID, skuID)
	if err != nil {
		if err.Error() == "record not found" {
			return domain.Inventory{}, errors.New("record not found")
		}
		logger.Error("Database error while fetching inventory", logrus.Fields{"hub_id": hubID, "sku_id": skuID, "error": err.Error()})
		return domain.Inventory{}, err
	}
	return inventory, nil
}
