package book

import (
	"github.com/gin-gonic/gin"
)

// update book's information, return after update book
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "updating book %s", id)
}
