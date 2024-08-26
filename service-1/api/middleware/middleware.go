package middleware

import (
    "auth/api/token"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }

        // Remove 'Bearer ' prefix if present
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        valid, err := token.ValidateToken(tokenString)
        if err != nil || !valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
            c.Abort()
            return
        }

        claims, err := token.ExtractClaim(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims", "details": err.Error()})
            c.Abort()
            return
        }
        c.Set("claims", claims)
        c.Next()
    }
}