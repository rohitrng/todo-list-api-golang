package middleware

import (
	"net/http"
	"todo-list-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.Jwtkey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired Token"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
