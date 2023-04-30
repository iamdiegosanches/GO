package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// books slice to seed book data.
var books = []Book{
	{ID: uuid.New(), Title: "Blue Train", Author: "John Coltrane", PublicationDate: time.Date(1957, 9, 15, 0, 0, 0, 0, time.UTC), Publisher: "Blue Note"},
	{ID: uuid.New(), Title: "To Kill a Mockingbird", Author: "Harper Lee", PublicationDate: time.Date(1960, 7, 11, 0, 0, 0, 0, time.UTC), Publisher: "J. B. Lippincott & Co."},
	{ID: uuid.New(), Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", PublicationDate: time.Date(1925, 4, 10, 0, 0, 0, 0, time.UTC), Publisher: "Charles Scribner's Sons"},
}

func handleGetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func handlePostBooks(c *gin.Context) {
	var newBook Book

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add a new book to the slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func handleGetById(c *gin.Context) {
	uuidParam := c.Param("uuid")
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid uuid"})
		return
	}
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range books {
		if a.ID == uuid {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
