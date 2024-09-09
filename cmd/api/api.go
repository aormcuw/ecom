package api

import (
	"log"
	"net/http"

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
	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, r)
}
