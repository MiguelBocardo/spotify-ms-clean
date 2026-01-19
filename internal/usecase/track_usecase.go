package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/example/spotify-ms-clean/internal/domain"
	"github.com/example/spotify-ms-clean/internal/infra/repo"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, val interface{})
}

type TrackUseCase struct {
	repo  repo.TrackRepo
	cache Cache
}

func NewTrackUseCase(r repo.TrackRepo, c Cache) *TrackUseCase {
	return &TrackUseCase{repo: r, cache: c}
}

func (uc *TrackUseCase) ListTracks() ([]domain.Track, error) {
	const key = "tracks_all"
	if v, ok := uc.cache.Get(key); ok {
		if b, ok2 := v.([]byte); ok2 {
			var out []domain.Track
			_ = json.Unmarshal(b, &out)
			return out, nil
		}
	}
	res, err := uc.repo.FindAll()
	if err != nil { return nil, err }
	b, _ := json.Marshal(res)
	uc.cache.Set(key, b)
	return res, nil
}

func (uc *TrackUseCase) GetTrack(id string) (domain.Track, error) {
	key := fmt.Sprintf("track_%s", id)
	if v, ok := uc.cache.Get(key); ok {
		if b, ok2 := v.([]byte); ok2 {
			var out domain.Track
			_ = json.Unmarshal(b, &out)
			return out, nil
		}
	}
	t, err := uc.repo.FindByID(id)
	if err != nil { return domain.Track{}, err }
	b, _ := json.Marshal(t)
	uc.cache.Set(key, b)
	return t, nil
}
