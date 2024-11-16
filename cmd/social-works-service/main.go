package main

import (
	"log"
	"os"

	"github.com/huxleyberg/socialworks/internal/app"
	"github.com/huxleyberg/socialworks/internal/db"
)

func main() {
	postgresDbUrl := os.Getenv("POSTGRES_DB_URL")
	dbConn := db.ProvidePostgres(postgresDbUrl)
	a := app.New(dbConn)

	mux := a.Handler()
	log.Fatal(a.Run(mux))
}
