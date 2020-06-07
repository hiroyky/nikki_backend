package main

import (
	"github.com/hiroyky/nikki_backend/config"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hiroyky/nikki_backend/resolver"
)

const defaultPort = "8080"

func main() {
	if config.Env.DebugMode {
		log.Println("Debug mode")
	}

	srv := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+config.Env.Port, nil))
}
