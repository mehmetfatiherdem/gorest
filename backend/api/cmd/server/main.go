package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// get handler
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// post handler
func addAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums := append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

// get /{id} handler
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, val := range albums {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	
	router.GET("/api/albums", getAlbums)
	router.GET("/api/albums/:id", getAlbumByID)
	router.POST("/api/albums", addAlbum)

	router.Run("localhost:8080")
}
