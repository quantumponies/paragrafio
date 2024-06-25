package routers

import (
	"paragrafio/pkg/services"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateStandaloneRouter(providers *services.Providers) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(60 * time.Second))

	return router
}