package storages

import (
	"example/web-service-gin/app/models"
)

type Storage struct {
	Albums []models.Album
}

func NewAlbumStorage() *Storage {
	return &Storage{}
}

func (storage *Storage) ReturnAlbums() []models.Album {
	return storage.Albums
}

func (storage *Storage) ReturnAlbumByID(id string) models.Album {
	var result models.Album
	for _, alb := range storage.Albums {
		if alb.ID == id {
			result = alb
		}
	}
	return result
}

func (storage *Storage) AppendAlbum(alb models.Album) error {
	storage.Albums = append(storage.Albums, alb)
	return nil
}
