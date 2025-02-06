package service

import (
	"context"
	"errors"
	"wms/domain"
	"wms/logger"
	"wms/repo"

	"github.com/sirupsen/logrus"
)

type SKUService interface {
	GetSKUByID(ctx context.Context, skuID string) (*domain.SKU, error)
	FetchSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error)
	CreateSKU(ctx context.Context, sku *domain.SKU) (*domain.SKU, error)
}

type skuService struct {
	repo repo.SKURepository
}

func NewSKUService(r repo.SKURepository) SKUService {
	return &skuService{repo: r}
}

func (s *skuService) GetSKUByID(ctx context.Context, skuID string) (*domain.SKU, error) {
	sku, err := s.repo.GetSKUByID(ctx, skuID)
	if err != nil {
		logger.Error("Failed to fetch SKU by ID", logrus.Fields{
			"sku_id": skuID,
			"error":  err.Error(),
		})
		return nil, errors.New("failed to fetch SKU by ID: " + err.Error())
	}
	return sku, nil
}

func (s *skuService) FetchSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error) {
	if sellerID <= 0 {
		return nil, errors.New("invalid Seller ID")
	}

	skus, err := s.repo.GetSkuBySellerID(ctx, sellerID)
	if err != nil {
		logger.Error("Failed to fetch SKUs by Seller ID", logrus.Fields{
			"seller_id": sellerID,
			"error":     err.Error(),
		})
		return nil, err
	}

	return skus, nil
}

func (s *skuService) CreateSKU(ctx context.Context, sku *domain.SKU) (*domain.SKU, error) {
	if sku.SellerID <= 0 || sku.PPU <= 0 {
		return nil, errors.New("invalid SKU data")
	}

	createdSKU, err := s.repo.CreateSKU(ctx, sku)
	if err != nil {
		logger.Error("Failed to create SKU", logrus.Fields{
			"error": err.Error(),
		})
		return nil, err
	}

	return createdSKU, nil
}
