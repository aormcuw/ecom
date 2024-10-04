package cart

import (
	"fmt"

	"github.com/aormcuw/ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity of product")
		}
		productIDs[i] = item.ProductID
	}
	return productIDs, nil
}

func (h *Handler) CreateOrder(ps []types.Product, items []types.CartItem, userID int64) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range ps {
		productMap[int(product.ID)] = product
	}
	// check if product is available
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}
	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// reduce quantitiy of product
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		if err := h.productStore.UpdateProduct(product); err != nil {
			return 0, 0, err
		}
	}
	// create order
	order := types.Order{
		UserID:  int(userID),
		Total:   totalPrice,
		Status:  "pending",
		Address: "unknown",
	}
	orderID, err := h.store.CreateOrder(order)
	if err != nil {
		return 0, 0, err
	}
	// create order items
	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[int]types.Product) error {
	if len(cartItems) == 0 {

		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {

		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d not found", item.ProductID)
		}
		if item.Quantity > product.Quantity {
			return fmt.Errorf("not enough stock of product %d", item.ProductID)
		}
	}

	return nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	total := 0.0
	for _, item := range items {
		product := productMap[item.ProductID]
		total += float64(item.Quantity) * product.Price
	}
	return total
}
