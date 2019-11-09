package book

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "got account %s", id)
}
func CreateBook(c *gin.Context) {
	fmt.Println("111")

	c.JSON(200, "creating account")
}
func UpdateBook(c *gin.Context) {
	fmt.Println("111")

	id := c.Param("id")
	c.String(200, "updating account %s", id)
}
func DeleteBook(c *gin.Context) {
	fmt.Println("111")
	id := c.Param("id")
	c.String(200, "deleting account %s", id)
	fmt.Println("111")
}
