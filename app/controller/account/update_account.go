package account

import (
	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	id := c.Param("id")
	c.String(200, "updating account %s", id)
}
