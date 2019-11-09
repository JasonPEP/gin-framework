package book

import (
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "got book %s", id)
}
func CreateBook(c *gin.Context) {
	c.String(200, "creating book")
}
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "updating book %s", id)
}
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "deleting book %s", id)
}
