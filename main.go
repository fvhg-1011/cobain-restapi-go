package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// initiate struct
type album struct {
	ID     string  `json:"id`
	Title  string  `json:"title`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// create slice froma struct album
var albums = []album{
	{ID: "1", Title: "Power of Love", Artist: "Jimi Hendrix", Price: 96.89},
	{ID: "2", Title: "Voodoo Child", Artist: "SVR", Price: 95},
	{ID: "3", Title: "Who Did You Think I Was", Artist: "John Mayer", Price: 99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsbyID)
	router.Run("localhost:8080")
}

// GET request
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST rqeust (add new album to exisitng album data)
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	//add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GET request from specific ID
func getAlbumsbyID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album"})
}
