package book

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-q-assignment-1/model"
	"net/http"
)

func Update(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)

		var book model.Book
		result := db.QueryRow(`SELECT * FROM book WHERE ID = ?`, id)

		err := result.Scan(&book.ID, &book.Title, &book.Code, &book.Author, &book.PublishYear, &book.Country)
		if err != nil {
			fmt.Printf("Error occured while load book data: %v\n", err)
			context.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Book not found",
			})
			return
		}

		if err = context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		_, err = db.Exec("UPDATE book SET title = ?, code = ?, author = ?, publishYear = ?, country = ? WHERE ID = ?", book.Title, book.Code, book.Author, book.PublishYear, book.Country, book.ID)
		if err != nil {
			fmt.Printf("Error occured while update book: %v\n", err)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	}
}
