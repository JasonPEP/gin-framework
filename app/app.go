package app

import (
	"fmt"
	middleware "gin-framework/app/middlewares"
	router "gin-framework/app/routers"
	"gin-framework/app/tools"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Engine *gin.Engine
)

func init() {
	fmt.Println("Application initializing")
	Engine = gin.New()
}
func Run() (err error) {
	// init middlewares
	middleware.InitMiddleware(Engine)
	// init routers
	router.InitRouters(Engine)
	fmt.Println(tools.TimeToTimestemp(time.Now()))
	return Engine.Run(":8080")
}
