package book

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Delete(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)

		remove, _ := db.Exec(`DELETE FROM book WHERE ID = ?`, id)
		count, err := remove.RowsAffected()
		if err != nil || count == 0 {
			fmt.Printf("Error occured while delete book: %v\n", err)
			context.JSON(http.StatusNotFound, gin.H{
				"status":  false,
				"message": "Book not found",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": true,
		})
	}
}
