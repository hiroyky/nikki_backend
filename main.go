package main

import (
	"github.com/hiroyky/nikki_backend/config"
	"github.com/hiroyky/nikki_backend/infrastructures"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	if config.Env.DebugMode {
		log.Println("Debug mode")
	}

	db, err := infrastructures.OpenDatabase(
		config.Env.DatabaseHostName,
		config.Env.DatabasePort,
		config.Env.DatabaseDBName,
		config.Env.DatabaseUserName,
		config.Env.DatabasePassword,
		config.Env.DebugMode,
	)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	srv := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+config.Env.Port, nil))
}
