package app

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/huxleyberg/socialworks/internal/config"
	"github.com/huxleyberg/socialworks/internal/env"
	"github.com/huxleyberg/socialworks/internal/health"
	"gorm.io/gorm"
)

type App struct {
	Config        config.Config
	HealthHandler health.HealthHandler
}

func New(postgresDBProvider *gorm.DB) App {
	cfg := config.Config{
		Addr: env.GetString("ADDR", ":8080"),
	}
	return App{Config: cfg}
}

func (a App) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.HealthHandler.HealthCheck)
	})

	return r
}

func (a App) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         a.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("server has started at %s", a.Config.Addr)
	return srv.ListenAndServe()
}
