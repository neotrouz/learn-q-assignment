package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"learn-q-assignment-1/model"
	"net/http"
)

func Delete(db *scribble.Driver) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)

		var book model.Book
		err := db.Read("book", id, &book)
		if err != nil {
			fmt.Println("Book not found", err)
			context.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Book not found",
			})
			return
		}

		err = db.Delete("book", id)
		if err != nil {
			fmt.Println("Error occured while delete book", err)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": err == nil,
		})
	}
}
