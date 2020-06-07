package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	DatabaseHostName string `envconfig:"DATABASE_HOSTNAME"`
	DatabasePort     string `envconfig:"DATABASE_PORT"`
	DatabaseUserName string `envconfig:"DATABASE_USERNAME"`
	DatabaseDBName   string `envconfig:"DATABASE_DB_NAME"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`
	DebugMode        bool   `envconfig:"DEBUG_MODE" default:"false"`
	Port             string `envconfig:"PORT"`
}

var Env environment

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := envconfig.Process("", &Env); err != nil {
		panic(err)
	}
}
