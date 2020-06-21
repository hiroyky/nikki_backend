package interfaces

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	r := gin.Default()

	admin := r.Group("/admin")
	{
		admin.GET("/", func(ctx *gin.Context) {
			playground.Handler("Admin graphql", "/admin/query").ServeHTTP(ctx.Writer, ctx.Request)
		})
		admin.POST("/query")
	}

	viewer := r.Group("/viewer")
	{
		viewer.GET("/", func(ctx *gin.Context) {
			playground.Handler("Viewer graphql", "/viewer/query").ServeHTTP(ctx.Writer, ctx.Request)
		})
		viewer.POST("/query")
	}

	Engine = r
}
