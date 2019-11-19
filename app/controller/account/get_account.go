package account

import (
	model "gin-framework/app/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAccount(c *gin.Context) {
	id := c.Param("id")

	account, e := model.GetAccount(id)
	if e != nil {
		log.Println(e)
		_ = c.AbortWithError(http.StatusNotFound, e)
		return
	}
	c.JSON(200, account)
}
