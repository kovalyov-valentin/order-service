package service

import (
	"context"
	"github.com/kovalyov-valentin/orders-service/internal/repository/cache"
	"github.com/kovalyov-valentin/orders-service/internal/models"
	"github.com/kovalyov-valentin/orders-service/internal/repository"
)


type OrderServicer interface {
	Create(ctx context.Context, orderUID string, order models.Order) error
	GetOrderByUID(ctx context.Context, uid string) (models.Order, error)

}

type Service struct {
	Repo OrderServicer
	Cache *cache.Cache
	ServiceOrder *OrderService

}

func NewService(repo *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Repo: repo,
		Cache: cache,
		ServiceOrder: NewOrderService(repo.OrderServicer, cache),
	}
}