package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroyky/nikki_backend/config"
)

func AdminBasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		config.Env.AdminUserName: config.Env.AdminPassword,
	})
}
