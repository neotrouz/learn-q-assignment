package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	scribble "github.com/nanobox-io/golang-scribble"
	"learn-q-assignment-1/model"
	"net/http"
	"strconv"
)

func Create(db *scribble.Driver) gin.HandlerFunc {
	return func(context *gin.Context) {
		index := model.Index{}
		err := db.Read("index", "data", &index)

		index.LastId = index.LastId + 1

		var book model.Book
		book.ID = index.LastId
		if err = context.ShouldBind(&book); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		err = db.Write("book", strconv.Itoa(index.LastId), book)
		if err != nil {
			fmt.Println("Error occured while create book", err)
		}

		err = db.Write("index", "data", index)
		if err != nil {
			fmt.Println("Error occured while save index", err)
		}
		context.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   book,
		})
	}
}
