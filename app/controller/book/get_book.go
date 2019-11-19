package book

import (
	model "gin-framework/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// return a book's detail information
func GetBook(c *gin.Context) {
	id := c.Param("id")

	book, e := model.GetBook(id)
	if e != nil {
		log.Println(e)
		_ = c.AbortWithError(http.StatusNotFound, e)
		return
	}
	c.JSON(200, book)
}
