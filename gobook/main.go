package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	ConnectDB()

	router.GET("/books", handleGetBooks)
	router.GET("/books/:uuid", handleGetById)
	router.POST("/books", handlePostBooks)
	router.DELETE("/books/:uuid", handleDeleteBook)

	router.Run(":8080")
}
