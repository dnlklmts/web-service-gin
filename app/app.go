package app

import (
	"example/web-service-gin/api"
	"example/web-service-gin/app/handlers"
	"example/web-service-gin/app/models"
	"example/web-service-gin/app/processors"
	"example/web-service-gin/app/storages"
)

var albums []models.Album

func init() {
	// albums slice to seed record album data.
	albums = []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

func Serve() {
	storage := storages.NewAlbumStorage()
	storage.Albums = albums

	processor := processors.NewAlbumProcessor(storage)
	handler := handlers.NewAlbumHandler(processor)
	router := api.CreateRouter(handler)

	router.Run("localhost:8080")
}
