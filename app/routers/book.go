package router

import (
	"gin-framework/app/controller/book"
	"github.com/gin-gonic/gin"
)

func InitBookRouter(engine *gin.Engine) {
	engine.GET("/book/:id", book.GetBook)
	engine.POST("/book", book.CreateBook)
	engine.PUT("/book/:id", book.UpdateBook)
	engine.DELETE("/book/:id", book.DeleteBook)
}
