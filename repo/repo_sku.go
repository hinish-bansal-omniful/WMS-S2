package repo

import (
	"context"
	"errors"
	"sync"
	"wms/domain"
	"wms/logger"

	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SKURepository interface {
	GetSKUByID(ctx context.Context, skuID string) (*domain.SKU, error)
	GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error)
	CreateSKU(ctx context.Context, sku *domain.SKU) (*domain.SKU, error)
}

type skuRepository struct {
	db *postgres.DbCluster
}

var skuRepo *skuRepository
var skuRepoOnce sync.Once

func NewSKURepository(db *postgres.DbCluster) SKURepository {
	skuRepoOnce.Do(func() {
		skuRepo = &skuRepository{db: db}
	})
	return skuRepo
}

func (r *skuRepository) GetSKUByID(ctx context.Context, skuID string) (*domain.SKU, error) {
	var sku domain.SKU
	err := r.db.GetMasterDB(ctx).Table("sku").Where("id = ?", skuID).First(&sku).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Info("SKU not found in database", logrus.Fields{
			"sku_id": skuID,
		})
		return nil, nil
	}

	if err != nil {
		logger.Error("Database error while fetching SKU", logrus.Fields{
			"sku_id": skuID,
			"error":  err.Error(),
		})
		return nil, err
	}

	return &sku, nil
}

func (r *skuRepository) GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.SKU, error) {
	var sku []domain.SKU

	if sellerID <= 0 {
		return nil, errors.New("invalid Seller ID")
	}

	err := r.db.GetMasterDB(ctx).Where("seller_id = ?", sellerID).Find(&sku).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Info("No SKUs found for seller", logrus.Fields{
			"seller_id": sellerID,
		})
		return nil, nil
	}

	if err != nil {
		logger.Error("Database error while fetching SKUs", logrus.Fields{
			"seller_id": sellerID,
			"error":     err.Error(),
		})
		return nil, err
	}

	return sku, nil
}

func (r *skuRepository) CreateSKU(ctx context.Context, sku *domain.SKU) (*domain.SKU, error) {
	if err := r.db.GetMasterDB(ctx).Create(&sku).Error; err != nil {
		logger.Error("Database error while creating SKU", logrus.Fields{
			"error": err.Error(),
		})
		return nil, err
	}
	return sku, nil
}
