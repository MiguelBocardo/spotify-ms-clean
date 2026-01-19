package repo

import "github.com/example/spotify-ms-clean/internal/domain"

type TrackRepo interface {
	FindAll() ([]domain.Track, error)
	FindByID(id string) (domain.Track, error)
}

type inMemoryTrackRepo struct {
	data map[string]domain.Track
}

func NewInMemoryTrackRepo() TrackRepo {
	return &inMemoryTrackRepo{
		data: map[string]domain.Track{
			"1": {ID: "1", Title: "Song A", Artist: "Artist X", Length: 210000},
			"2": {ID: "2", Title: "Song B", Artist: "Artist Y", Length: 185000},
			"3": {ID: "3", Title: "Song C", Artist: "Artist Z", Length: 240000},
		},
	}
}

func (r *inMemoryTrackRepo) FindAll() ([]domain.Track, error) {
	res := make([]domain.Track, 0, len(r.data))
	for _, v := range r.data {
		res = append(res, v)
	}
	return res, nil
}

func (r *inMemoryTrackRepo) FindByID(id string) (domain.Track, error) {
	if t, ok := r.data[id]; ok {
		return t, nil
	}
	return domain.Track{}, nil
}
