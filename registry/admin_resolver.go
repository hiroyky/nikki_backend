package registry

import "github.com/hiroyky/nikki_backend/resolvers/adminresolvers"

var AdminResolverConfig adminresolvers.Config

func init() {
	AdminResolverConfig = adminresolvers.Config{}
	AdminResolverConfig.Resolvers = &adminresolvers.Resolver{}
}
