package registry

import (
	"github.com/hiroyky/nikki_backend/resolvers/viewerresolvers"
)

var ViewerResolverConfig viewerresolvers.Config

func init() {
	ViewerResolverConfig = viewerresolvers.Config{}
	ViewerResolverConfig.Complexity = *viewerresolvers.ViewerComplexRoot
	ViewerResolverConfig.Resolvers = &viewerresolvers.Resolver{}
}
