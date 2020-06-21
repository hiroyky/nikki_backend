package main

import (
	"github.com/hiroyky/nikki_backend/config"
	"github.com/hiroyky/nikki_backend/infrastructures"
	"github.com/hiroyky/nikki_backend/interfaces"
	"log"
	"net/http"
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

	if err := interfaces.Engine.Run(); err != nil {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Env.Port)
	log.Fatal(http.ListenAndServe(":"+config.Env.Port, nil))
}
