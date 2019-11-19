package book

import (
	entity "gin-framework/app/entities"
	bookService "gin-framework/app/service/book"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type creatorPayload struct {
	Title     string `form:"title" binding:"required"`
	AccountId int64  `form:"account_id" binding:"required"`
}

// create a new book, return new book information
func CreateBook(c *gin.Context) {
	payload := creatorPayload{}
	e := c.ShouldBind(&payload)
	if e != nil {
		log.Println(e)
	}
	book, e := bookService.NewBook(payload.Title, payload.AccountId)
	if e != nil {
		_ = c.AbortWithError(http.StatusNotFound, e)
		return
	}
	c.JSON(http.StatusOK, entity.Result{
		Code: 0,
		Data: book,
		Msg:  "ok",
	})
}
