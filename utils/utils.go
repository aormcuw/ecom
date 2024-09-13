package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckAndBind(c *gin.Context, payload any) {
	if c.Request.Body == nil || c.Request.ContentLength == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is empty"})
		return
	}
	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body",
			"details": err.Error()})
	}
}
