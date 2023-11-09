package service

import (
	"context"
	"log"

	"github.com/kovalyov-valentin/orders-service/internal/repository/cache"
	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/kovalyov-valentin/orders-service/internal/repository"
)

type OrderService struct {
	repo  repository.OrderServicer
	cache cache.Cache
}

func NewOrderService(repo repository.OrderServicer, cache *cache.Cache) *OrderService {
	return &OrderService{
		repo:  repo,
		cache: *cache,
	}
}

func (s *OrderService) Create(ctx context.Context, orderUID string, order models.Order) error {
	s.cache.Set(&order)
	return s.repo.Create(ctx, orderUID, order)
}

func (s *OrderService) GetOrderByUID(ctx context.Context, uid string) (models.Order, error) {
	data, err := s.cache.Get(uid)
	if err != nil {
		log.Println("error receiving cached data", err)
	}

	if data != nil {
		return *data, nil
	}

	return s.repo.GetOrderByUID(ctx, uid)
}

