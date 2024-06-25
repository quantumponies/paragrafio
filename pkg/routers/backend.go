package routers

import (
	"net/http"
	"paragrafio/pkg/services"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
)

func prepareCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("web.allowed_origins"),
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}

func indexRouter(providers *services.Providers) http.Handler {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeHTML))
	router.Use(prepareCORS().Handler)
	
	

	return router
}

func CreateBackendRouter(providers *services.Providers) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Compress(5))
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/", indexRouter(providers))

	return router
}