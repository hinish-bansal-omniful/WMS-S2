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

type Repository interface {
	GetAllHubs(ctx context.Context) ([]domain.Hub, error)
	GetHubByID(ctx context.Context, id int) (domain.Hub, error)        // Added this method
	CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) // **New**
}

type repository struct {
	db *postgres.DbCluster
}

var repo *repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) Repository {
	repoOnce.Do(func() {
		repo = &repository{db: db}
	})
	return repo
}

func (r *repository) GetAllHubs(ctx context.Context) ([]domain.Hub, error) {
	var hubs []domain.Hub
	err := r.db.GetMasterDB(ctx).Find(&hubs).Error
	if err != nil {
		logger.Error("Failed to fetch hubs from database", logrus.Fields{
			"error": err.Error(),
			"query": "SELECT * FROM hubs",
		})
		return nil, err
	}
	return hubs, nil
}

func (r *repository) GetHubByID(ctx context.Context, id int) (domain.Hub, error) {
	var hub domain.Hub
	if id <= 0 {
		return hub, errors.New("invalid ID")
	}

	result := r.db.GetMasterDB(ctx).First(&hub, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return domain.Hub{}, errors.New("record not found") // Explicit error handling
		}
		logger.Error("Error retrieving hub from database", logrus.Fields{"hub_id": id, "error": result.Error.Error()})
		return domain.Hub{}, result.Error
	}

	return hub, nil
}

func (r *repository) CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) {

	db := r.db.GetMasterDB(ctx)
	result := db.Create(&hub)
	if result.Error != nil {
		logger.Error("Failed to insert hub", logrus.Fields{"error": result.Error.Error()})
		return domain.Hub{}, result.Error
	}

	return hub, nil
}
