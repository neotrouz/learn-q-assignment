package main

import (
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"learn-q-assignment-1/controllers/book"
	"log"
	"net/http"
)

var (
	engine *gin.Engine
	db     *scribble.Driver
)

func init() {
	engine = gin.Default()

	var err error
	db, err = scribble.New("db", nil)
	if err != nil {
		log.Fatalln("Error occured connect DB", err)
	}
}

func main() {
	listingHandler()

	err := engine.Run(":8081")
	if err != nil {
		log.Fatalln("Error occured running server", err)
	}
}

func listingHandler() {
	// Create Book
	engine.POST("/book", book.Create(db))

	// List Book
	engine.GET("/book", book.List(db))

	// Detail Book
	engine.GET("/book/:id", book.Detail(db))

	// Update Book
	engine.PATCH("/book/:id", book.Update(db))

	// Delete Book
	engine.DELETE("/book/:id", book.Delete(db))

	// 404
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Page Not Found",
		})
	})
}
