package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	s := NewAPIServer(":8080")
	router.GET("/books", makeHTTPHandleFunc(s.handleGetBook))
	router.POST("/books", makeHTTPHandleFunc(s.handleCreateBook))

	router.Run()
}
