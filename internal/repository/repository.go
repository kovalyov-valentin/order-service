package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/kovalyov-valentin/orders-service/internal/models"
)

type OrderServicer interface {
	Create(ctx context.Context, orderUID string, order models.Order) error
	GetOrderByUID(ctx context.Context, uid string) (models.Order, error)
}

type Repository struct {
	OrderServicer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		OrderServicer: NewOrderPostgres(db),
	}
}
