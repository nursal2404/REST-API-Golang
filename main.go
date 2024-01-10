package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nursal2404/REST-API-Golang"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("books/:id/:title", booksHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBooksHandler)

	router.Run(":8080")
}