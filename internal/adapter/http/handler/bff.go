package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/example/spotify-ms-clean/internal/usecase"
)

type BFFHandler struct{ uc *usecase.TrackUseCase }

func NewBFFHandler(uc *usecase.TrackUseCase) *BFFHandler { return &BFFHandler{uc: uc} }

// AggregatedTracks simula agregação (ex.: metadados + popularidade + disponibilidade)
func (h *BFFHandler) AggregatedTracks(w http.ResponseWriter, r *http.Request) {
	tracks, err := h.uc.ListTracks()
	if err != nil { http.Error(w, err.Error(), 500); return }
	writeJSON(w, http.StatusOK, map[string]interface{}{"items": tracks, "count": len(tracks)})
}

func (h *BFFHandler) GetTrack(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	t, err := h.uc.GetTrack(id)
	if err != nil { http.Error(w, err.Error(), 500); return }
	if t.ID == "" { http.NotFound(w, r); return }
	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
