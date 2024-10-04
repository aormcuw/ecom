package cart

import (
	"net/http"

	"github.com/aormcuw/ecom/service/auth"

	"github.com/aormcuw/ecom/types"
	"github.com/aormcuw/ecom/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
	userStore    types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore))
}

func (h *Handler) handleCheckout(c *gin.Context) {
	userID := auth.GetUserIDFromContext(c)
	var cart types.CartCheckPayload
	// Bind the incoming JSON to the `cart` struct
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.Validate.Struct(&cart); err != nil { //+
		if validationErrors, ok := err.(validator.ValidationErrors); ok { //+
			var errors []string                  //+
			for _, v := range validationErrors { //+
				errors = append(errors, v.Field()+": "+v.Tag()) //+
			} //+
			c.JSON(400, gin.H{ //+
				"error":   "Invalid request body", //+
				"details": errors,                 //+
			}) //+
			return //+
		} //+
	} //+

	// get product
	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := h.productStore.GetProductByIDs(productIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	orderID, totalPrice, err := h.CreateOrder(ps, cart.Items, int64(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"order_id": orderID, "total_price": totalPrice})

}
