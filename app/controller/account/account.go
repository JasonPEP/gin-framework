package account

import (
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	id := c.Param("id")

	c.String(200, "got account %s", id)
}
func CreateAccount(c *gin.Context) {
	c.JSON(200, "creating account")
}
func UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "updating account %s", id)
}
func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "deleting account %s", id)
}
