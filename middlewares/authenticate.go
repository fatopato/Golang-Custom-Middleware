package middlewares

import (
	"net/http"
	"strings"

	"github.com/fatopato/custom-middlewares/security"
	"github.com/gin-gonic/gin"
)


func AuthenticationMiddleware(c *gin.Context) {
	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Split the header to get the token part
	tokenString := strings.Split(authHeader, "Bearer ")[1]

	_, err := security.ValidateJWTToken(tokenString)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Token is valid, proceed with the request
	c.Next()
}