package registry

import (
	"github.com/hiroyky/nikki_backend/resolvers/viewerresolvers"
)

var ViewerResolverConfig viewerresolvers.Config

func init() {
	ViewerResolverConfig = viewerresolvers.Config{}
	ViewerResolverConfig.Complexity = newViewerComplexRoot()
	ViewerResolverConfig.Resolvers = &viewerresolvers.Resolver{}
}

func newViewerComplexRoot() viewerresolvers.ComplexityRoot {
	root := viewerresolvers.ComplexityRoot{}
	root.Query.Articles = viewerresolvers.ArticlesComplexity
	root.Query.Article = viewerresolvers.ArticleComplexity
	return root
}
