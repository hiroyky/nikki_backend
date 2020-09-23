package interfaces

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/hiroyky/nikki_backend/config"
	"github.com/hiroyky/nikki_backend/interfaces/middleware"
	"github.com/hiroyky/nikki_backend/registry"
	"github.com/hiroyky/nikki_backend/resolvers/adminresolvers"
	"github.com/hiroyky/nikki_backend/resolvers/viewerresolvers"
)

var Engine *gin.Engine
var adminServer *handler.Server
var viewerServer *handler.Server

func init() {
	r := gin.Default()
	r.Use(middleware.Cors())
	adminServer = handler.NewDefaultServer(adminresolvers.NewExecutableSchema(registry.AdminResolverConfig))
	viewerServer = handler.NewDefaultServer(viewerresolvers.NewExecutableSchema(registry.ViewerResolverConfig))
	viewerServer.Use(extension.FixedComplexityLimit(config.GraphQLComplexityOfViewer))

	admin := r.Group("/admin", middleware.AdminBasicAuth())
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
