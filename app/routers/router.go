package router

import (
	"fmt"

	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	fmt.Println("Routers initializing")
}
func InitRouters(engine *gin.Engine) {
	InitAccountRouter(engine)
	InitBookRouter(engine)
	engine.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
}
