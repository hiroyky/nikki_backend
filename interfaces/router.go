package interfaces

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hiroyky/nikki_backend/resolvers/adminresolvers"
	"github.com/hiroyky/nikki_backend/resolvers/viewerresolvers"
)

var Engine *gin.Engine
var adminServer *handler.Server
var viewerServer *handler.Server

func init() {
	r := gin.Default()
	adminServer = handler.NewDefaultServer(adminresolvers.NewExecutableSchema(adminresolvers.Config{Resolvers: &adminresolvers.Resolver{}}))
	viewerServer = handler.NewDefaultServer(viewerresolvers.NewExecutableSchema(viewerresolvers.Config{Resolvers: &viewerresolvers.Resolver{}}))

	admin := r.Group("/admin")
	{
		admin.GET("/", func(ctx *gin.Context) {
			playground.Handler("Admin graphql", "/admin/query").ServeHTTP(ctx.Writer, ctx.Request)
		})
		admin.POST("/query", func(ctx *gin.Context) {
			adminServer.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	viewer := r.Group("/viewer")
	{
		viewer.GET("/", func(ctx *gin.Context) {
			playground.Handler("Viewer graphql", "/viewer/query").ServeHTTP(ctx.Writer, ctx.Request)
		})
		viewer.POST("/query", func(ctx *gin.Context) {
			viewerServer.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}
	Engine = r
}
