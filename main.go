package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"learn-q-assignment-1/controllers/book"
	"log"
	"net/http"
)

var (
	engine *gin.Engine
	//db     *scribble.Driver
	db *sql.DB
)

func init() {
	engine = gin.Default()

	var err error
	db, err = sql.Open("mysql", "root@/learn-q_golang")
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
