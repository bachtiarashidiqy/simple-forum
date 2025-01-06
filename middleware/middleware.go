package middleware

import (
	"fmt"
	"github.com/bachtiarashidiqy/simple-forum/internal/configs"
	"github.com/bachtiarashidiqy/simple-forum/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Authorization header is empty"))
		}
		userID, username, err := jwt.ValidateToken(secretKey, header)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}
		c.Set("UserID", userID)
		c.Set("Username", username)
		c.Next()
	}
}
