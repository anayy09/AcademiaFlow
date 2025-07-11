package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:5173"} // SvelteKit dev servers
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
    config.ExposeHeaders = []string{"Content-Length"}
    config.AllowCredentials = true

    return cors.New(config)
}