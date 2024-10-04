package order

import (
	"github.com/aormcuw/ecom/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	// Create the order using GORM's Create method
	if err := s.db.Create(&order).Error; err != nil {
		return 0, err
	}

	// Return the ID of the newly created order and no error
	return int(order.ID), nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	// Use GORM to insert the new order item
	if err := s.db.Create(&orderItem).Error; err != nil {
		return err
	}

	return nil
}
