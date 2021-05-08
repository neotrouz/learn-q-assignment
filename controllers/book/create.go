package book

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-q-assignment-1/model"
	"net/http"
)

func Create(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var book model.Book
		if err := context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		insert, err := db.Exec(`INSERT INTO book (title, code, author, publishYear, country) VALUES (?, ?, ?, ?, ?)`, book.Title, book.Code, book.Author, book.PublishYear, book.Country)
		if err != nil {
			fmt.Printf("Error occured while create book: %v\n", err)
		}

		lastID, _ := insert.LastInsertId()
		book.ID = int(lastID)

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	}
}
