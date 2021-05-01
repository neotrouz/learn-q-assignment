package book

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"learn-q-assignment-1/model"
	"net/http"
)

func List(db *scribble.Driver) gin.HandlerFunc {
	return func(context *gin.Context) {
		records, err := db.ReadAll("book")
		if err != nil {
			fmt.Println("Error occured while get list books", err)
		}

		var books []model.Book
		for _, f := range records {
			var foundBook model.Book
			err = json.Unmarshal([]byte(f), &foundBook)
			if err != nil {
				fmt.Println("Error occured while loop", err)
			}
			books = append(books, foundBook)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   books,
		})
	}
}
