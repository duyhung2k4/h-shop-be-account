package config

import (
	"flag"
	"log"
)

func init() {
	var migrate bool = false

	flag.BoolVar(&migrate, "db", true, "Migrate Database?")

	loadEnv()

	if err := connectPostgresql(migrate); err != nil {
		log.Fatalf("Error connect Postgresql: %v", err)
	}
	connectRedis()
}
