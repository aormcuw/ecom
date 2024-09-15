package user

import (
	"github.com/aormcuw/ecom/service/auth"
	"github.com/aormcuw/ecom/types"
	"github.com/aormcuw/ecom/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/login", h.HandleLogin)
	r.POST("/register", h.HandleRegister)
}

func (h *Handler) HandleLogin(c *gin.Context) {
	// implement login logic here
	c.JSON(200, gin.H{"message": "login successful"})
}

func (h *Handler) HandleRegister(c *gin.Context) {
	var payload types.RegisterUserPayload //+

	if err := c.ShouldBindJSON(&payload); err != nil { //+
		c.JSON(400, gin.H{"error": err.Error()}) //+
		return                                   //+
	} //+

	// validate the payload
	if err := utils.Validate.Struct(&payload); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, v := range validationErrors {
				errors = append(errors, v.Field()+": "+v.Tag())
			}
			c.JSON(400, gin.H{
				"error":   "Invalid request body",
				"details": errors,
			})
			return
		}
	}

	// check if the user is already registered
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}
	//+
	// Hash password before storing
	hashedPassword, err := auth.HashedPassword(payload.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//+
	// if not, create a new user in the database
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//+
	// if everything is successful, return a success message
	c.JSON(200, gin.H{"message": "registration successful"})
}
