package api

import (
	"example/web-service-gin/app/handlers"

	"github.com/gin-gonic/gin"
)

func CreateRouter(handler *handlers.AlbumHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.POST("/albums", handler.PostAlbum)
	router.GET("/albums/:id", handler.GetAlbumByID)

	return router
}
