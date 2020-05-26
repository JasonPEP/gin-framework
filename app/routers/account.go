package router

import (
	"gin-framework/app/controller/account"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitAccountRouter(engine *gin.Engine) {
	engine.GET("account/:id", account.GetAccount)
	engine.POST("account", account.CreateAccount)
	engine.PUT("account/:id", account.UpdateAccount)
	engine.DELETE("account/:id", account.DeleteAccount)
	engine.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}
