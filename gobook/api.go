package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//connection := "mongodb+srv://admin:admin@cluster0.5ljbs7k.mongodb.net/?retryWrites=true&w=majority"
	router := gin.Default()

	router.GET("/books", handleGetBooks)
	router.GET("/books/:uuid", handleGetById)
	router.POST("/books", handlePostBooks)

	router.Run()
}
