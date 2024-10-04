package products

import (
	"fmt"
	"strings"

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

func (s *Store) GetProductByIDs(productionIDs []int) ([]types.Product, error) {
	placeholders := strings.Repeat(",?", len(productionIDs)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholders)

	// convert productsid to interface{}
	args := make([]interface{}, len(productionIDs))
	for i, v := range productionIDs {
		args[i] = v
	}

	// Create a slice to store the products
	var products []types.Product

	// Execute the query using GORM's Raw() and scan the results into products
	if err := s.db.Raw(query, args...).Scan(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	if err := s.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
