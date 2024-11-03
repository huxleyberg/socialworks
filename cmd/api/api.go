package main

import "net/http"

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {
	srv := &http.Server{
		Addr: app.config.addr,
	}

	return srv.ListenAndServe()
}
