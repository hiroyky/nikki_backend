package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiroyky/nikki_backend/config"
	"strings"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     strings.Split(config.Env.CorsAllowOrigins, ","),
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept", "Content-Length", "Cookie", "Referer", "Access-Control-Allow-Credentials", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Type", "Content-Length", "Set-Cookie", "X-Maintenance-Status"},
		AllowCredentials: true,
		MaxAge:           -1,
	})
}
