package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowMethods = append(config.AllowMethods, "DELETE", "OPTIONS")
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	config.AllowCredentials = false

	return cors.New(config)
}
