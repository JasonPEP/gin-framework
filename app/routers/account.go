package router

import (
	"gin-framework/app/controller/account"

	"github.com/gin-gonic/gin"
)

func InitAccountRouter(engine *gin.Engine) {
	engine.GET("account/:id", account.GetAccount)
	engine.POST("account", account.CreateAccount)
	engine.PUT("account/:id", account.UpdateAccount)
	engine.DELETE("account/:id", account.DeleteAccount)
}
