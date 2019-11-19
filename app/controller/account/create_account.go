package account

import (
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	c.String(200, "creating account")
}
