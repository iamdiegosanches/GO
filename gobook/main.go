package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	s := NewAPIServer(":8080")
	router.GET("/books", func(c *gin.Context) {
		err := s.handleGetBook(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	router.Run()
}
