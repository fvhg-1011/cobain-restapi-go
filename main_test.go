package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetAlbums tests the GET /albums endpoint.
func TestGetAlbums(t *testing.T) {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetAlbumsByID tests the GET /albums/:id endpoint.
func TestGetAlbumsByID(t *testing.T) {
	router := gin.Default()
	router.GET("/albums/:id", getAlbumsbyID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestPostAlbums tests the POST /albums endpoint.
func TestPostAlbums(t *testing.T) {
	router := gin.Default()
	router.POST("/albums", postAlbums)

	// Create a request body.
	jsonData := `{"id":"4", "title":"New Album", "artist":"New Artist", "price":100.0}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer([]byte(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

