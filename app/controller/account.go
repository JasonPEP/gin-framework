package controller

import (
	"github.com/gin-gonic/gin"
)

func GetAccount(ctx *gin.Context) {
	ctx.JSON(200, "account")
}
