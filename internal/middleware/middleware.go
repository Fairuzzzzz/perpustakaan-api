package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Fairuzzzzz/perpustakaan-api/internal/configs"
	"github.com/Fairuzzzzz/perpustakaan-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware() gin.HandlerFunc {
	secretJWT := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			log.Error().Msg("Missing token in request")
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, role, username, err := jwt.ValidateToken(header, secretJWT)
		if err != nil {
			log.Error().Err(err).Msg("Failed to validate token")
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Set("role", role)
		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			log.Error().Msg("Role not found in context")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to access this resource",
			})
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			log.Error().Msg("Role is not a string")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid role format",
			})
			return
		}

		if roleStr != "admin" {
			log.Error().Str("role", roleStr).Msg("Non-admin role attempting to access admin resource")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to access this resource",
			})
			return
		}

		c.Next()
	}
}
