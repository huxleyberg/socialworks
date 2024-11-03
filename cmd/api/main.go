package main

import (
	"log"

	"github.com/huxleyberg/socialworks/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.routes()
	log.Fatal(app.run(mux))
}
