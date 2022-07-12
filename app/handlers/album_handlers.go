package handlers

import (
	"example/web-service-gin/app/models"
	"example/web-service-gin/app/processors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	processor *processors.AlbumProcessor
}

func NewAlbumHandler(processor *processors.AlbumProcessor) *AlbumHandler {
	return &AlbumHandler{processor: processor}
}

func (handler *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := handler.processor.ListAlbums()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (handler *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	alb, err := handler.processor.FindAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, alb)
}

func (handler *AlbumHandler) PostAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}

	if err := handler.processor.CreateAlbum(newAlbum); err != nil {
		c.JSON(http.StatusBadRequest,
			map[string]string{
				"result": "error",
				"error":  err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, newAlbum)
}
