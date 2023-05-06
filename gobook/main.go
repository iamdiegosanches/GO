package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	client := ConnectDB()

	router.GET("/books", func(c *gin.Context) {
		handleGetBooks(c, client)
	})
	router.GET("/books/:uuid", func(c *gin.Context) {
		handleGetById(c, client)
	})
	router.POST("/books", func(c *gin.Context) {
		handlePostBooks(c, client)
	})
	router.DELETE("/books/:uuid", func(c *gin.Context) {
		handleDeleteBook(c, client)
	})

	router.Run(":8080")
}
