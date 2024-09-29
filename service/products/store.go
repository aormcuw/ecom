package products

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

func (s *Store) GetProducts() ([]types.Product, error) {
	var products []types.Product

	err := s.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
