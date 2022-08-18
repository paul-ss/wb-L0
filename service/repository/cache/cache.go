package cache

import (
	"github.com/paul-ss/wb-L0/service/domain"
	"github.com/paul-ss/wb-L0/service/repository/postgres"
)

var (
	cache *Cache
)

// Cache TODO threadsafe??
type Cache struct {
	pg domain.Repository
	m  map[string][]byte
}

func NewCache() *Cache {
	if cache != nil {
		return cache
	}

	cache = &Cache{
		pg: postgres.NewPgConn(),
		m:  make(map[string][]byte),
	}

	return cache
}

func (c *Cache) StoreOrder(id string, order []byte) error {
	err := c.pg.StoreOrder(id, order)
	if err != nil {
		return err
	}

	c.m[id] = order
	return nil
}

func (c *Cache) GetOrderById(id string) ([]byte, error) {
	order, ok := c.m[id]
	if ok {
		return order, nil
	}

	order, err := c.pg.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	c.m[id] = order
	return order, nil
}
