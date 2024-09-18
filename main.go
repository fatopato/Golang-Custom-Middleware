package main

import (
	"net/http"

	"github.com/fatopato/custom-middlewares/middlewares"
	"github.com/fatopato/custom-middlewares/services"

	"github.com/fatopato/custom-middlewares/security"

	"github.com/fatopato/custom-middlewares/entity"
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()

    // Public endpoints
    r.POST("/register", register)
    r.POST("/login", login)

    // Protected endpoints
    protected := r.Group("/api")
    protected.Use(middlewares.AuthenticationMiddleware)
    
    protected.GET("/topSecret", topSecret)

    r.Run()
}

func register(c *gin.Context) {

    var user entity.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    err := services.RegisterUser(user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func login(c *gin.Context) {
    var user entity.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }


    err := services.ValidateUser(user)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }


    tokenString, err := security.CreateJWTToken(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
        return
    }


    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func topSecret(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "This is a top secret message"})
}
