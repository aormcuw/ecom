package api

import (
	"log"
	"net/http"

	"github.com/aormcuw/ecom/service/cart"
	"github.com/aormcuw/ecom/service/order"
	"github.com/aormcuw/ecom/service/products"
	"github.com/aormcuw/ecom/service/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := gin.Default()
	subrouter := r.Group("/api/v1/")
	subrouter.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "GET /api/v1/users"})
	})
	userStore := user.NewStore(s.db) // Initialize user store with the provided DB connection
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := products.NewStore(s.db) // Initialize product store with the
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, r)
}
