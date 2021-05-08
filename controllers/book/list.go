package book

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-q-assignment-1/model"
	"net/http"
)

func List(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		result, err := db.Query(`SELECT * FROM book`)
		if result == nil || err != nil {
			fmt.Printf("Error occured while get list books: %v\n", err)
			context.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  false,
				"message": "Error occured while get list books",
			})
			return
		}

		var books []model.Book
		for result.Next() {
			var foundBook model.Book
			err = result.Scan(&foundBook.ID, &foundBook.Title, &foundBook.Code, &foundBook.Author, &foundBook.PublishYear, &foundBook.Country)
			if err != nil {
				fmt.Printf("Error occured while loop: %v\n", err)
			}
			books = append(books, foundBook)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   books,
		})
	}
}
