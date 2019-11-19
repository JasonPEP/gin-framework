package account

import (
	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "deleting account %s", id)
}
