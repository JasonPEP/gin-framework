package book

import (
	"github.com/gin-gonic/gin"
)

// delete a book, return bool
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "deleting book %s", id)
}
