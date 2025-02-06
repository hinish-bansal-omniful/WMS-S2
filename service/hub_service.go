package service

import (
	"context"
	"errors"
	"wms/domain"
	"wms/logger"
	"wms/repo"

	"github.com/sirupsen/logrus"
)

type HubService interface {
	FetchHubs(ctx context.Context) ([]domain.Hub, error)
	FetchHubByID(ctx context.Context, id int) (domain.Hub, error)
	CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) // **New**
}

type Service struct {
	repo repo.Repository
}

func NewService(r repo.Repository) HubService {
	return &Service{repo: r}
}

func (s *Service) FetchHubs(ctx context.Context) ([]domain.Hub, error) {
	hubs, err := s.repo.GetAllHubs(ctx)
	if err != nil {
		logger.Error("Failed to fetch hubs from repository", logrus.Fields{"error": err.Error()})
		return nil, errors.New("failed to fetch hubs from repository: " + err.Error())
	}
	return hubs, nil
}

func (s *Service) FetchHubByID(ctx context.Context, id int) (domain.Hub, error) {
	hub, err := s.repo.GetHubByID(ctx, id)
	if err != nil {
		if err.Error() == "record not found" { // Check for specific DB errors
			return domain.Hub{}, errors.New("record not found")
		}
		logger.Error("Database error while fetching hub", logrus.Fields{"hub_id": id, "error": err.Error()})
		return domain.Hub{}, err
	}
	return hub, nil
}

func (s *Service) CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) {
	// Basic validation
	if hub.Name == "" || hub.Location == "" {
		return domain.Hub{}, errors.New("hub name and location are required")
	}

	createdHub, err := s.repo.CreateHub(ctx, hub)
	if err != nil {
		logger.Error("Failed to create hub in repository", logrus.Fields{"error": err.Error()})
		return domain.Hub{}, err
	}

	return createdHub, nil
}
