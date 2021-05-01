package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"learn-q-assignment-1/model"
	"net/http"
)

func Update(db *scribble.Driver) gin.HandlerFunc {
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

		if err = context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		err = db.Write("book", id, book)
		if err != nil {
			fmt.Println("Error occured while update book", err)
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	}
}
