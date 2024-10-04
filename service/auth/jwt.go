package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aormcuw/ecom/config"
	"github.com/aormcuw/ecom/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
)

func CreateJWT(secret []byte, UserID string) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTDurationinSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": UserID,
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WithJWTAuth(handlerFunc gin.HandlerFunc, store types.UserStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the JWT token from the Authorization header
		tokenString := getTokenFromRequest(c)

		// Validate the token
		token, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Fetch the userID from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// Safely assert and convert userID from claims
		var userID int
		switch v := claims["userID"].(type) {
		case string:
			userID, err = strconv.Atoi(v)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userID in token"})
				return
			}
		case float64:
			userID = int(v) // JWT often stores numbers as float64
		default:
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userID in token"})
			return
		}

		// Fetch the user from the database using the userID
		u, err := store.GetUserByIds(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Set context "userID" to the userID
		c.Set("userID", userID)

		// Attach the user object to the context so the handler can access it
		c.Set("user", u)

		// Call the next handler function
		handlerFunc(c)
	}
}

func getTokenFromRequest(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, _ := ctx.Value("userID").(int)
	return userID
}
