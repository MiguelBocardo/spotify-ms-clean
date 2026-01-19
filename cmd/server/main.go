package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/example/spotify-ms-clean/internal/adapter/http/handler"
	"github.com/example/spotify-ms-clean/internal/infra/cache"
	"github.com/example/spotify-ms-clean/internal/infra/repo"
	"github.com/example/spotify-ms-clean/internal/usecase"
)

func main() {
	port := getenv("PORT", "8080")

	repo := repo.NewInMemoryTrackRepo()
	c := cache.NewInMemoryCache(5 * time.Second)

	uc := usecase.NewTrackUseCase(repo, c)
	bff := handler.NewBFFHandler(uc)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
		r.Get("/tracks", bff.AggregatedTracks) // BFF: agrega múltiplas fontes
		r.Get("/tracks/{id}", bff.GetTrack)
	})

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
