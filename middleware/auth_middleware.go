package middleware

import (
	"fmt"
	"main/config/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{
				"error": "token tidak ditemukan",
			})
			c.Abort()
			return
		}
		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(401, gin.H{
				"error": "format token salah",
			})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		token = strings.TrimSpace(token)

		fmt.Println("Authorization:", auth)
		fmt.Println("Token:", token)

		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "token tidak valid",
			})
			c.Abort()
			return
		}
		// simpan claims ke context
		c.Set("claims", claims)
		// ganti request dengan context yang baru
		c.Next()
	}
}
