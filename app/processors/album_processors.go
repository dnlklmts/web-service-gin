package processors

import (
	"errors"
	"example/web-service-gin/app/models"
	"example/web-service-gin/app/storages"
)

type AlbumProcessor struct {
	storage *storages.Storage
}

func NewAlbumProcessor(storage *storages.Storage) *AlbumProcessor {
	return &AlbumProcessor{storage: storage}
}

func (proc *AlbumProcessor) ListAlbums() ([]models.Album, error) {
	return proc.storage.ReturnAlbums(), nil
}

func (proc *AlbumProcessor) FindAlbumByID(id string) (models.Album, error) {
	alb := proc.storage.ReturnAlbumByID(id)
	if alb.ID != id {
		return alb, errors.New("album not found")
	}
	return alb, nil
}

func (proc *AlbumProcessor) CreateAlbum(alb models.Album) error {
	return proc.storage.AppendAlbum(alb)
}
