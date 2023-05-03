package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRuter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestPostBookHandler(t *testing.T) {
	router := SetUpRuter()
	router.POST("/books", handlePostBooks)
	book := NewBook("Test", "Test Author", time.Now(), "Test Publisher")
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetBooksHandler(t *testing.T) {
	router := SetUpRuter()
	router.GET("/books", handleGetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteBooksHandler(t *testing.T) {

}
