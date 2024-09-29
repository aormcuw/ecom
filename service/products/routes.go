package products

import (
	"github.com/aormcuw/ecom/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/products", h.handleCreateProduct)
}

func (h *Handler) handleCreateProduct(c *gin.Context) {
	ps, err := h.store.GetProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.ShouldBind(ps)
	c.JSON(200, gin.H{
		"message": "Product list retrieved successfully",
	})

}
