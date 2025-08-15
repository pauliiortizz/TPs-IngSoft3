package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

// Claims estructura para los claims del token JWT
type Claims struct {
	Username string `json:"username"`
	Tipo     bool   `json:"tipo"`
	jwt.StandardClaims
}

// AuthMiddleware es un middleware para autenticación
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("tipo", claims.Tipo)
		c.Next()
	}
}

// AdminMiddleware verifica que el usuario es un administrador
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tipo, exists := c.Get("tipo")
		if !exists || !tipo.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// StudentMiddleware verifica que el usuario es un estudiante
func StudentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tipo, exists := c.Get("tipo")
		if !exists || tipo.(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Student access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
