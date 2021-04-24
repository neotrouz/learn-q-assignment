package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Code   string `json:"code"`
	Author string `json:"author"`
}

type Index struct {
	LastId int
}

func main() {
	engine := gin.Default()

	db, err := scribble.New("db", nil)
	if err != nil {
		log.Fatalln("Error occured connect DB", err)
	}

	listingHandler(engine, db)

	err = engine.Run(":8081")
	if err != nil {
		log.Fatalln("Error occured running server", err)
	}
}

func listingHandler(engine *gin.Engine, db *scribble.Driver) {
	engine.POST("/book", func(context *gin.Context) {
		index := Index{}
		err := db.Read("book", "index", &index)
		if err != nil {
			log.Fatalln("Error occured while load index", err)
		}

		index.LastId = index.LastId + 1
		book := Book{
			ID:     index.LastId,
			Title:  context.PostForm("title"),
			Code:   context.PostForm("code"),
			Author: context.PostForm("author"),
		}

		err = db.Write("book", strconv.Itoa(index.LastId), book)
		if err != nil {
			log.Fatalln("Error occured while create book", err)
		}

		err = db.Write("book", "index", index)
		if err != nil {
			log.Fatalln("Error occured while save index", err)
		}
		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	})

	engine.GET("/book", func(context *gin.Context) {
		records, err := db.ReadAll("book")
		if err != nil {
			log.Fatalln("Error occured while get list books", err)
		}

		books := []Book{}
		for _, f := range records {
			foundBook := Book{}
			err = json.Unmarshal([]byte(f), &foundBook)
			if err != nil {
				log.Fatalln("Error occured while loop", err)
			}
			books = append(books, foundBook)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   books,
		})
	})

	engine.GET("/book/:id", func(context *gin.Context) {
		fmt.Println(context.Param("id"))
		book := Book{}
		err := db.Read("book", context.Param("id"), &book)
		if err != nil {
			log.Fatalln("Error occured while get book", err)
		}
		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	})

	engine.PUT("/book/:id", func(context *gin.Context) {
		id := context.Param("id")

		book := Book{}
		err := db.Read("book", id, &book)
		if err != nil {
			log.Fatalln("Error occured while get book", err)
		}

		book = Book{
			Title:  context.PostForm("title"),
			Code:   context.PostForm("code"),
			Author: context.PostForm("author"),
		}

		err = db.Write("book", id, book)
		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	})

	engine.DELETE("/book/:id", func(context *gin.Context) {
		id := context.Param("id")

		book := Book{}
		err := db.Read("book", id, &book)
		if err != nil {
			log.Fatalln("Error occured while delete book", err)
		}
		err = db.Delete("book", id)
		if err != nil {
			log.Fatalln("Error occured while delete book", err)
		}
		context.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	})

	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Page Not Found",
		})
	})
}
