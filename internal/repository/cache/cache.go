package cache

import (
	"errors"
	"sync"

	"github.com/kovalyov-valentin/orders-service/internal/models"
)

type Cache struct {
	sync.RWMutex
	Order        map[string]*models.Order

}

func NewCache() *Cache {
	return &Cache{
		Order: make(map[string]*models.Order),
	}
}

func (c *Cache) Set(order *models.Order) error {
	c.Lock()
	defer c.Unlock()

	c.Order[order.OrderUID] = order

	return nil
}

func (c *Cache) Get(key string) (*models.Order, error) {
	c.RLock()
	defer c.RUnlock()
	data, ex := c.Order[key]

	if !ex {
		return nil, errors.New("no such element in memory cache")
	}

	return data, nil
}


