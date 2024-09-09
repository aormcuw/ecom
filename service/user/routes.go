package user

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHanler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/login", h.HandleLogin)
	r.POST("/register", h.HandleRegister)
}

func (h *Handler) HandleLogin(c *gin.Context) {
	// implement login logic here
	c.JSON(200, gin.H{"message": "login successful"})
}

func (h *Handler) HandleRegister(c *gin.Context) {
	// implement registration logic here
	c.JSON(200, gin.H{"message": "registration successful"})
}
